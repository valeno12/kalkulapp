package session

import (
	"context"
	"fmt"

	db "github.com/valeno12/kalkulapp/internal/models"
	"github.com/valeno12/kalkulapp/internal/logger"
)

func (s *Service) JoinSession(ctx context.Context, code string, userName string) (int64, error) {
	session, err := s.queries.GetSessionByCode(ctx, code)
	if err != nil {
		logger.Log.Warnf("Sesión con código %s no encontrada", code)
		return 0, err
	}

	if session.MaxParticipants.Valid {
		count, err := s.queries.CountUsersInSession(ctx, session.ID)
		if err != nil {
			logger.Log.Error("Error al contar usuarios en sesión:", err)
			return 0, err
		}

		if count >= int(session.MaxParticipants.Int32) {
			logger.Log.Warnf("Sesión %s llena, no se pueden unir más usuarios", code)
			return 0, fmt.Errorf("la sesión está llena")
		}
	}

	result, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		SessionID: session.ID,
		Name:      userName,
	})
	if err != nil {
		logger.Log.Error("Error al insertar usuario en la sesión:", err)
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		logger.Log.Error("Error al obtener el ID del usuario insertado:", err)
		return 0, err
	}

	logger.Log.Infof("Usuario %s agregado a la sesión %s con ID %d", userName, code, userID)
	return userID, nil
}
