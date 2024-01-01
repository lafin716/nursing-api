package main

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/pkg/bootstrap"
	"nursing_api/pkg/configs"
	"nursing_api/pkg/middlewares"
	"nursing_api/pkg/routes"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// fiber 생성
	config := configs.FiberConfig()
	app := fiber.New(config)

	// 미들웨어 설정
	middlewares.FiberMiddleware(app)

	// 전체 허용 API 라우팅 설정
	routes.PublicRoutes(app)
	// 인증 필요 API 라우팅 설정
	routes.PrivateRoutes(app)
	// 존재하지 않는 엔드포인트 라우팅 설정
	routes.NotFoundRoute(app)

	// 서버 환경에 따라 서버 구동 분기
	if os.Getenv("PROFILE") == "production" {
		bootstrap.StartServerWithGracefulShutdown(app)
	} else {
		bootstrap.StartServer(app)
	}
}
