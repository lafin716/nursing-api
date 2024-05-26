package main

import (
	"nursing_api/internal/config"
	"nursing_api/pkg/database"
)

func main() {
	databaseConfig := config.NewDatabaseConfig()
	databaseClient := database.NewPostgresClient(databaseConfig)
	databaseClient.Migrate()
}
