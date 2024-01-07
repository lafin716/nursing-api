package database

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"log"
	"nursing_api/pkg/ent"
	"os"
)

var Set = wire.NewSet(NewPostgresClient)

type Config struct {
	Host      string
	Port      int
	Database  string
	User      string
	Password  string
	SSLEnable bool
	Debug     bool
}

type DatabaseClient struct {
	Client *ent.Client
	Ctx    context.Context
}

func (c *Config) Serialize() string {
	sslMode := "disable"
	if c.SSLEnable {
		sslMode = "require"
	}

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Database,
		c.Password,
		sslMode,
	)
}

func NewPostgresClient(config *Config) *DatabaseClient {
	log.Printf("데이터베이스 URL : %s", config.Serialize())
	client, err := ent.Open("postgres", config.Serialize())
	if err != nil {
		log.Printf("데이터베이스 연결에 실패하였습니다: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Printf("데이터베이스 컨텍스트 생성에 실패하였습니다: %v", err)
	}

	return &DatabaseClient{
		Client: client,
		Ctx:    ctx,
	}
}

func (d *DatabaseClient) Close() {
	d.Client.Close()
}

// 마이그레이션
func (d *DatabaseClient) Migrate() {
	if err := d.Client.Schema.WriteTo(d.Ctx, os.Stdout); err != nil {
		log.Fatalf("마이그레이션 수행 중 오류 발생: %v", err)
	}
}
