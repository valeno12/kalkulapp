package session

import (
	"context"
	"database/sql"
	"errors"

	"github.com/valeno12/kalkulapp/internal/logger"
)

func (s *SessionService) LeaveSession(ctx context.Context, code string, userID int64) error {
	// Buscar la sesión por código
	session, err := s.queries.GetSessionByCode(ctx, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("código de sesión inválido")
		}
		logger.Log.Error("Error al buscar sesión:", err)
		return err
	}

	logger.Log.Infof("Validando salida - UserID: %d, Anfitrión: %d, Sesión: %s", userID, session.CreatedBy, code)

	// Verificar si el usuario pertenece a la sesión
	user, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("el usuario no existe en esta sesión")
		}
		logger.Log.Error("Error al obtener usuario:", err)
		return err
	}

	if user.SessionID != session.ID {
		return errors.New("el usuario no pertenece a esta sesión")
	}

	if userID == session.CreatedBy {
		logger.Log.Warnf("Intento de salida del anfitrión (ID: %d) en sesión %s", userID, code)
		return errors.New("el anfitrión no puede abandonar la sesión, debe eliminarla")
	}

	// Eliminar usuario de la sesión
	err = s.queries.DeleteUser(ctx, userID)
	if err != nil {
		logger.Log.Error("Error al eliminar usuario de la sesión:", err)
		return err
	}

	logger.Log.Infof("Usuario %d salió de la sesión %s", userID, code)
	return nil
}
