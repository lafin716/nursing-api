package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"nursing_api/pkg/bootstrap"
	"nursing_api/pkg/configs"
	"nursing_api/pkg/database"
	"os"
)

func main() {

	// DB 객체 생성
	log.Println("서버 구동 시작")

	log.Println("데이터베이스 연결 시작")
	dbConfig := configs.DatabaseConfig()
	client, err := database.NewPostgresClient(dbConfig)
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	log.Println("데이터베이스 연결 완료")

	// 마이그레이션
	log.Println("데이터베이스 마이그레이션 체크")
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("마이그레이션 수행 중 오류 발생: %v", err)
	}
	log.Println("데이터베이스 마이그레이션 완료")

	// fiber 생성
	log.Println("프레임워크 구동 시작")
	config := configs.FiberConfig()
	app := fiber.New(config)

	// 서버 구동 시작
	server := bootstrap.NewServer(app, client)
	server.Run()
}
