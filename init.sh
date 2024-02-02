#!/bin/bash

echo ".env 파일 검사"
ENV_FILE=".env"
if [ ! -e $ENV_FILE ]; then
  cp .env.example .env
  chmod 744 .env
  echo ".env 파일이 신규 생성되었습니다. 로컬 환경에 맞게 변수값을 설정하세요."
  exit 1
fi

echo "DB 스키마 초기화"
go generate ./pkg/ent

echo "의존성 초기화"
wire ./internal

echo "API 명세서 생성"
swag init -g cmd/main.go