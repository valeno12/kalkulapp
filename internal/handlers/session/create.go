package session

import (
	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/logger"
	"github.com/valeno12/kalkulapp/internal/utils"
)

func (h *SessionHandler) CreateSession(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.CreateSessionRequest
	if err := c.Bind(&req); err != nil {
		logger.Log.Warnf("Error al parsear request en CreateSession: %v", err)
		return utils.ErrorResponse(c, "Formato de datos incorrecto")
	}

	// Validaciones básicas
	if req.SessionName == "" {
		return utils.ErrorResponse(c, "El nombre de la sesión es obligatorio")
	}

	if req.UserName == "" {
		return utils.ErrorResponse(c, "El nombre del usuario creador es obligatorio")
	}

	sessionID, code, err := h.service.CreateSession(ctx, req)
	if err != nil {
		logger.Log.Errorf("Error al crear sesión: %v", err)
		return utils.ErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, "Sesión creada correctamente", dto.CreateSessionResponse{
		SessionID: sessionID,
		Code:      code,
	})
}
