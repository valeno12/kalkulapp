package session

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/utils"
)

func (h *SessionHandler) LeaveSession(c echo.Context) error {
	ctx := c.Request().Context()
	code := c.Param("code")

	if code == "" {
		return utils.ErrorResponse(c, "El código de sesión es obligatorio")
	}

	// Obtener el ID del usuario que quiere salir (desde Query Param)
	userIDStr := c.QueryParam("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, "ID de usuario inválido")
	}

	// Intentar eliminar al usuario
	err = h.service.LeaveSession(ctx, code, userID)
	if err != nil {
		return utils.ErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, "El usuario ha salido de la sesión", nil)
}
