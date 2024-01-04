package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"nursing_api/app/bootstrap"
	"nursing_api/pkg/server"
	"os"
)

func main() {
	log.Println("서버 구동 시작")

	// 서버 구동 시작
	fiberApp := bootstrap.GetServer()
	serverConfig := &server.ServerConfig{
		Profile: os.Getenv("PROFILE"),
		App:     fiberApp,
	}
	s := server.NewServer(serverConfig)
	err := s.Run()
	if err != nil {
		log.Println("서버 종료.")
	}
}
