package bootstrap

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"nursing_api/ent"
	"nursing_api/pkg/middlewares"
	"nursing_api/pkg/routes"
	"os"
	"os/signal"
)

type Server struct {
	app       *fiber.App
	entClient *ent.Client
}

func NewServer(app *fiber.App, c *ent.Client) *Server {
	return &Server{
		app:       app,
		entClient: c,
	}
}

func (s *Server) Run() error {
	// 미들웨어 설정
	log.Println("미들웨어 설정")
	middlewares.FiberMiddleware(s.app)

	// 전체 허용 API 라우팅 설정
	log.Println("미들웨어 설정")
	routes.PublicRoutes(s.app)

	// 인증 필요 API 라우팅 설정
	log.Println("퍼블릭 라우팅 설정")
	routes.PrivateRoutes(s.app)

	// 존재하지 않는 엔드포인트 라우팅 설정
	log.Println("프라이빗 라우팅 설정")
	routes.NotFoundRoute(s.app)

	// 서버 환경에 따라 서버 구동 분기
	if os.Getenv("PROFILE") == "production" {
		log.Println("운영 환경에서 서버를 구동합니다")
		StartServerWithGracefulShutdown(s.app)
	} else {
		log.Println("개발 환경에서 서버를 구동합니다")
		StartServer(s.app)
	}

	return nil
}

func StartServer(app *fiber.App) {
	fiberUrl := createConnectUrl()
	if err := app.Listen(fiberUrl); err != nil {
		log.Printf("서버 구동 중 에러 발생 : %v", err)
	}
}

func StartServerWithGracefulShutdown(app *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("서버 중지 중 에러 발생 : %v", err)
		}

		close(idleConnsClosed)
	}()

	fiberUrl := createConnectUrl()
	if err := app.Listen(fiberUrl); err != nil {
		log.Printf("서버 구동 중 에러 발생 : %v", err)
	}

	<-idleConnsClosed
}

func createConnectUrl() string {
	log.Printf("서버 포트 : %s", os.Getenv("PORT"))
	return fmt.Sprintf(
		":%s",
		os.Getenv("PORT"),
	)
}
