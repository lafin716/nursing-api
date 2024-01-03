package repository

import (
	"log"
	"nursing_api/ent"
	"nursing_api/pkg/database"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("데이터베이스 포트 형식이 아닙니다: %v", err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return database.NewPostgresClient(config)
}
