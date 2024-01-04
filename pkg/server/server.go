package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

type ServerConfig struct {
	Profile string
	App     *fiber.App
}

type Server struct {
	config *ServerConfig
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() error {
	// 서버 환경에 따라 서버 구동 분기
	if s.config.Profile == "production" {
		log.Println("운영 환경에서 서버를 구동합니다")
		StartServerWithGracefulShutdown(s.config.App)
	} else {
		log.Println("개발 환경에서 서버를 구동합니다")
		StartServer(s.config.App)
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
