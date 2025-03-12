package handlers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/logger"
	"github.com/valeno12/kalkulapp/internal/services"
	"github.com/valeno12/kalkulapp/internal/utils"
)

type SessionHandler struct {
	service *services.SessionService
}

func NewSessionHandler(service *services.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) CreateSession(c echo.Context) error {
	var req dto.CreateSessionRequest
	if err := c.Bind(&req); err != nil {
		logger.Log.Warnf("Error al parsear request: %v", err)
		return utils.ErrorResponse(c, "Datos inválidos")
	}

	sessionID, code, err := h.service.CreateSession(context.Background(), req.SessionName, req.UserName, req.MaxParticipants)
	if err != nil {
		logger.Log.Errorf("Error al crear sesión: %v", err)
		return utils.ErrorResponse(c, "No se pudo crear la sesión")
	}

	logger.Log.Infof("Sesión creada con éxito: ID %d, Código %s", sessionID, code)
	return utils.SuccessResponse(c, "Sesión creada correctamente", dto.CreateSessionResponse{
		SessionID: sessionID,
		Code:      code,
	})
}
