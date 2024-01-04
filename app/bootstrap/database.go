package bootstrap

import (
	"log"
	"nursing_api/pkg/database"
	"os"
	"strconv"
)

func GetDatabaseClient() *database.DatabaseClient {
	client := database.NewPostgresClient(getDatabaseConfig())
	return client
}

func getDatabaseConfig() *database.Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("데이터베이스 포트 형식이 아닙니다: %v", err)
	}

	sslMode, err := strconv.ParseBool(os.Getenv("DB_SSL"))
	if err != nil {
		sslMode = false
	}

	dbDebug, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if err != nil {
		dbDebug = false
	}

	config := &database.Config{
		Host:      os.Getenv("DB_HOST"),
		Port:      port,
		Database:  os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		SSLEnable: sslMode,
		Debug:     dbDebug,
	}

	return config
}
