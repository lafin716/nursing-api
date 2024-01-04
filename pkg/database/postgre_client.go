package database

import (
	"context"
	"fmt"
	"log"
	"nursing_api/ent"
	"os"
)

type Config struct {
	Host      string
	Port      int
	Database  string
	User      string
	Password  string
	SSLEnable string
}

type DatabaseClient struct {
	Client *ent.Client
	Ctx    context.Context
}

func (c *Config) Serialize() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		c.Host,
		c.Port,
		c.User,
		c.Database,
		c.Password)
}

func NewPostgresClient(config *Config) *DatabaseClient {
	client, err := ent.Open("postgres", config.Serialize())
	if err != nil {
		log.Printf("데이터베이스 연결에 실패하였습니다: %v", err)
	}

	defer client.Close()
	ctx := context.Background()

	return &DatabaseClient{
		Client: client,
		Ctx:    ctx,
	}
}

// 마이그레이션
func (d *DatabaseClient) Migrate() {
	if err := d.Client.Schema.WriteTo(d.Ctx, os.Stdout); err != nil {
		log.Fatalf("마이그레이션 수행 중 오류 발생: %v", err)
	}
}
