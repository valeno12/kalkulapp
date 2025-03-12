package session

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/logger"
	"github.com/valeno12/kalkulapp/internal/services/session"
	"github.com/valeno12/kalkulapp/internal/utils"
)

type JoinSessionRequest struct {
	UserName string `json:"user_name" validate:"required"`
}

func JoinSessionHandler(service *session.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")

		var req JoinSessionRequest
		if err := c.Bind(&req); err != nil {
			logger.Log.Warn("Error al parsear el request en JoinSessionHandler")
			return utils.ErrorResponse(c, "Datos inválidos")
		}

		userID, err := service.JoinSession(context.Background(), code, req.UserName)
		if err != nil {
			logger.Log.Error("Error al unir usuario a la sesión: ", err)
			return utils.ErrorResponse(c, "No se pudo unir a la sesión")
		}

		logger.Log.Infof("Usuario %s unido a la sesión %s (ID: %d)", req.UserName, code, userID)
		return utils.SuccessResponse(c, "Usuario unido correctamente", map[string]int64{"user_id": userID})
	}
}
