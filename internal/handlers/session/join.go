package session

import (
	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/logger"
	"github.com/valeno12/kalkulapp/internal/utils"
)

func (h *SessionHandler) JoinSession(c echo.Context) error {
	ctx := c.Request().Context()
	code := c.Param("code")

	if code == "" {
		return utils.ErrorResponse(c, "El código de sesión es obligatorio")
	}

	var req dto.JoinSessionRequest
	if err := c.Bind(&req); err != nil {
		logger.Log.Warn("Error al parsear el request en JoinSession")
		return utils.ErrorResponse(c, "Datos inválidos")
	}

	if req.UserName == "" {
		return utils.ErrorResponse(c, "El nombre del usuario es obligatorio")
	}

	userID, err := h.service.JoinSession(ctx, code, req.UserName)
	if err != nil {
		return utils.ErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, "Usuario unido correctamente", map[string]int64{"user_id": userID})
}
