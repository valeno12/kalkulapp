package session

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/valeno12/kalkulapp/internal/logger"
	db "github.com/valeno12/kalkulapp/internal/models"
)

func (s *SessionService) JoinSession(ctx context.Context, code string, userName string) (int64, error) {
	// Buscar la sesión por código
	session, err := s.queries.GetSessionByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("código de sesión inválido")
		}
		logger.Log.Error("Error al buscar sesión:", err)
		return 0, err
	}

	// Obtener la lista de usuarios en la sesión
	users, err := s.queries.GetUsersBySessionID(ctx, session.ID)
	if err != nil {
		logger.Log.Error("Error al obtener participantes de la sesión:", err)
		return 0, err
	}

	// Validar si el nombre ya existe en esta sesión
	for _, user := range users {
		if user.Name == userName {
			return 0, errors.New("ya existe un usuario con ese nombre en la sesión")
		}
	}
	// Verificar límite de participantes
	if session.MaxParticipants.Valid {
		count, err := s.queries.CountUsersInSession(ctx, session.ID)
		if err != nil {
			logger.Log.Error("Error al contar usuarios en la sesión:", err)
			return 0, err
		}

		if int32(count) >= session.MaxParticipants.Int32 {
			return 0, fmt.Errorf("la sesión está llena")
		}
	}

	// Insertar usuario en la sesión
	result, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		SessionID: session.ID,
		Name:      userName,
	})
	if err != nil {
		logger.Log.Error("Error al insertar usuario en la sesión:", err)
		return 0, err
	}

	// Obtener el ID del usuario insertado
	userID, err := result.LastInsertId()
	if err != nil {
		logger.Log.Error("Error al obtener el ID del usuario insertado:", err)
		return 0, err
	}

	logger.Log.Infof("Usuario %s agregado a la sesión %s con ID %d", userName, code, userID)
	return userID, nil
}
