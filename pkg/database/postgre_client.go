package database

import (
	"context"
	"fmt"
	"log"
	"nursing_api/pkg/ent"
	"os"
)

type Config struct {
	Host      string
	Port      int
	Database  string
	User      string
	Password  string
	SSLEnable bool
	Debug     bool
}

type DatabaseClient interface {
	GetClient() *ent.Client
	GetCtx() context.Context
	RegisterTx(tx Transactional) error
	GetTxManager() TransactionManager
	Migrate()
}

type databaseClient struct {
	Client    *ent.Client
	Ctx       context.Context
	TxManager TransactionManager
}

func (d *databaseClient) GetClient() *ent.Client {
	if d.TxManager.IsRunningTransaction() {
		return d.TxManager.GetTx().Client()
	}

	return d.Client
}

func (d *databaseClient) GetCtx() context.Context {
	return d.Ctx
}

func (d *databaseClient) RegisterTx(tx Transactional) error {
	return d.TxManager.Register(tx)
}

func (d *databaseClient) GetTxManager() TransactionManager {
	return d.TxManager
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

// 마이그레이션
func (d *databaseClient) Migrate() {
	if err := d.Client.Schema.WriteTo(d.Ctx, os.Stdout); err != nil {
		log.Fatalf("마이그레이션 수행 중 오류 발생: %v", err)
	}
}

func NewPostgresClient(config *Config) DatabaseClient {
	log.Printf("데이터베이스 URL : %s", config.Serialize())
	client, err := ent.Open("postgres", config.Serialize())
	if err != nil {
		log.Printf("데이터베이스 연결에 실패하였습니다: %v", err)
	}

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Printf("데이터베이스 컨텍스트 생성에 실패하였습니다: %v", err)
	}

	return &databaseClient{
		Client:    client,
		Ctx:       ctx,
		TxManager: NewTransactionManager(client, ctx),
	}
}
