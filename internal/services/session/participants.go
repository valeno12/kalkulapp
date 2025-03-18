package session

import (
	"context"
	"database/sql"
	"errors"

	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/logger"
)

func (s *SessionService) GetSessionParticipants(ctx context.Context, code string) ([]dto.Participant, error) {
	// Buscar la sesión por código
	session, err := s.queries.GetSessionByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("código de sesión inválido")
		}
		logger.Log.Error("Error al buscar sesión:", err)
		return nil, err
	}

	// Obtener los usuarios de la sesión
	users, err := s.queries.GetUsersBySessionID(ctx, session.ID)
	if err != nil {
		logger.Log.Error("Error al obtener participantes de la sesión:", err)
		return nil, err
	}

	// Convertir la respuesta al formato DTO
	participants := make([]dto.Participant, len(users))
	for i, user := range users {
		participants[i] = dto.Participant{
			ID:   user.ID,
			Name: user.Name,
		}
	}

	return participants, nil
}
