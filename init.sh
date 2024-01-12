#!/bin/bash

echo "DB 스키마 초기화"
go generate ./pkg/ent

echo "의존성 초기화"
wire ./internal

echo "API 명세서 생성"
swag init -g cmd/main.go