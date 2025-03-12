package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/database"
	"github.com/valeno12/kalkulapp/internal/handlers"
	db "github.com/valeno12/kalkulapp/internal/models"
	"github.com/valeno12/kalkulapp/internal/services"
)

func RegisterRoutes(e *echo.Echo) {
	queries := db.New(database.DB)
	sessionService := services.NewSessionService(queries)
	sessionHandler := handlers.NewSessionHandler(sessionService)

	e.POST("/sessions", sessionHandler.CreateSession)
}
