package session

import (
	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/dto"
	"github.com/valeno12/kalkulapp/internal/utils"
)

func (h *SessionHandler) GetSessionParticipants(c echo.Context) error {
	ctx := c.Request().Context()
	code := c.Param("code")

	if code == "" {
		return utils.ErrorResponse(c, "El código de sesión es obligatorio")
	}

	participants, err := h.service.GetSessionParticipants(ctx, code)
	if err != nil {
		return utils.ErrorResponse(c, err.Error())
	}

	return utils.SuccessResponse(c, "Participantes de la sesión", dto.GetSessionParticipantsResponse{
		Participants: participants,
	})
}
