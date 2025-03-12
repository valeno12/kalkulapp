package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/valeno12/kalkulapp/internal/database"
	"github.com/valeno12/kalkulapp/internal/logger"
	"github.com/valeno12/kalkulapp/internal/routes"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Iniciando servidor.")

	database.InitDB()
	defer database.CloseDB()

	e := echo.New()

	routes.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log.Infof("Servidor iniciado en el puerto %s.", port)
	log.Fatal(e.Start(":" + port))
}
