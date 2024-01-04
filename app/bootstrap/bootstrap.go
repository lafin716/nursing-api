package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetServer() *fiber.App {
	log.Println("JWT 모듈 초기화")
	jwtClient := GetJwtClient()
	log.Println("데이터베이스 초기화")
	dbClient := GetDatabaseClient()
	log.Println("Fiber 프레임워크 구동")
	fiberContainer := NewFiberApplication(jwtClient, dbClient)

	return fiberContainer.container.App
}
