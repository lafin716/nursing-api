package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"nursing_api/internal/config"
	"nursing_api/pkg/database"
)

func main() {
	databaseConfig := config.NewDatabaseConfig()
	databaseClient := database.NewPostgresClient(databaseConfig)
	databaseClient.Migrate()
}
