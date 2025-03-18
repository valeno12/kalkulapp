package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/database"
	handler "github.com/valeno12/kalkulapp/internal/handlers/session" // Alias para handlers
	db "github.com/valeno12/kalkulapp/internal/models"
	service "github.com/valeno12/kalkulapp/internal/services/session" // Alias para servicios
)

func RegisterRoutes(e *echo.Echo) {
	queries := db.New(database.DB)

	sessionService := service.NewSessionService(database.DB, queries)

	sessionHandler := handler.NewSessionHandler(sessionService)

	e.POST("/sessions", sessionHandler.CreateSession)
	e.POST("/sessions/:code/join", sessionHandler.JoinSession)
	e.GET("/sessions/:code/participants", sessionHandler.GetSessionParticipants)
	e.DELETE("/sessions/:code/leave", sessionHandler.LeaveSession)
}
