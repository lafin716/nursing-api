# Nursing API
> NursingAPI는 간병관리를 위한 API 서버입니다. <br>
> `Golang`기반 웹 프레임워크인 `Fiber`로 제작되었으며, 아래 상세 내용을 참고해주세요.

## Table of Contents
- [Requirements](#Requirements)
- [Installation](#Installation)

## Requirements
- golang 1.21.x
- postgresql

## Installation
아래 명령어를 통해 API를 로컬에서 설치 및 실행할 수 있습니다.

```shell
git clone https://github.com/lafin716/nursing-api
cd nursing-api

# DB, DIP 초기화 수행
./init.sh 

# .env 생성 후 DB 설정값등을 설정 후 아래 명령어로 서버 실행
go run main.go
```

