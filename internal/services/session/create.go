package session

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/logger"
	db "github.com/valeno12/kalkulapp/internal/models"
)

func (s *SessionService) CreateSession(ctx context.Context, req dto.CreateSessionRequest) (int64, string, error) {

	code := generateSessionCode()

	// Crear sesión primero con CreatedBy = 0 (temporalmente)
	params := db.CreateSessionParams{
		Name:            req.SessionName,
		Code:            code,
		CreatedBy:       0, // Se actualizará después
		MaxParticipants: sqlNullInt32(req.MaxParticipants),
	}

	if req.MaxParticipants != nil && *req.MaxParticipants < 2 {
		logger.Log.Warnf("Intento de crear sesión con max_participants inválido: %d", *req.MaxParticipants)
		return 0, "", errors.New("El número máximo de participantes debe ser mayor a 1")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Error("Error al iniciar transacción:", err)
		return 0, "", err
	}
	defer tx.Rollback()

	qtx := s.queries.WithTx(tx)

	// Crear sesión
	sessionResult, err := qtx.CreateSession(ctx, params)
	if err != nil {
		logger.Log.Error("Error al crear la sesión:", err)
		return 0, "", err
	}

	sessionID, err := sessionResult.LastInsertId()
	if err != nil {
		logger.Log.Error("Error al obtener ID de la sesión:", err)
		return 0, "", err
	}

	// Insertar usuario creador con el session_id correcto
	userResult, err := qtx.CreateUser(ctx, db.CreateUserParams{
		SessionID: sessionID,
		Name:      req.UserName,
	})
	if err != nil {
		logger.Log.Error("Error al agregar usuario a la sesión:", err)
		return 0, "", err
	}

	userID, err := userResult.LastInsertId()
	if err != nil {
		logger.Log.Error("Error al obtener ID del usuario creador:", err)
		return 0, "", err
	}

	// Actualizar el CreatedBy en la sesión con el userID
	err = qtx.UpdateSessionCreatedBy(ctx, db.UpdateSessionCreatedByParams{
		CreatedBy: userID,
		ID:        sessionID,
	})
	if err != nil {
		logger.Log.Error("Error al actualizar el anfitrión de la sesión:", err)
		return 0, "", err
	}

	// Confirmar transacción
	if err := tx.Commit(); err != nil {
		logger.Log.Error("Error al confirmar transacción:", err)
		return 0, "", err
	}

	logger.Log.Infof("Sesión creada exitosamente: ID %d, Código %s, Anfitrión %d", sessionID, code, userID)
	return sessionID, code, nil
}

// Manejo de valores nulos
func sqlNullInt32(val *int) sql.NullInt32 {
	if val == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: int32(*val), Valid: true}
}

// Generador de código optimizado
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateSessionCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}
