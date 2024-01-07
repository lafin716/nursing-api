package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"nursing_api/internal"
	"nursing_api/pkg/web"
	"os"
	"strings"
)

func main() {
	log.Println("서버 구동 시작")
	s, err := server.New()
	if err != nil {
		panic(err)
	}

	// 서버 환경에 따라 서버 구동 분기
	profile := os.Getenv("PROFILE")
	if strings.TrimSpace(profile) == "" {
		profile = "production"
	}
	if profile == "production" {
		log.Println("운영 환경에서 서버를 구동합니다")
		web.StartServerWithGracefulShutdown(s.App())
	} else {
		log.Println("개발 환경에서 서버를 구동합니다")
		web.StartServer(s.App())
	}

	web.StartServer(s.App())
	defer s.DB().Close()
}
