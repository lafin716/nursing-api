package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

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
