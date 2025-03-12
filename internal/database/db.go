package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/valeno12/kalkulapp/internal/logger"
)

var DB *sql.DB

func InitDB() {
	logger.Log.Info("Iniciando conexión a la base de datos.")

	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("Error cargando el archivo .env:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 1; i <= maxRetries; i++ {
		logger.Log.Infof("Intentando conectar a la base de datos (Intento %d/%d)...", i, maxRetries)
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			logger.Log.Info("Conexión a la base de datos establecida correctamente.")
			return
		}

		logger.Log.Warnf("Error al conectar a la base de datos: %v", err)
		time.Sleep(retryDelay)
	}

	logger.Log.Fatal("No se pudo conectar a la base de datos después de varios intentos.")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		logger.Log.Info("Conexión a la base de datos cerrada.")
	}
}
