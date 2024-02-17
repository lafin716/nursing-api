package database

import (
	"context"
	"errors"
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

type DatabaseClient struct {
	Client *ent.Client
	Ctx    context.Context
	Tx     *ent.Tx
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
		Tx:     nil,
	}
}

func (d *DatabaseClient) BeginTx() error {
	if d.Tx != nil {
		return errors.New("이미 실행중인 트랜잭션이 있습니다")
	}

	tx, err := d.Client.Tx(d.Ctx)
	if err != nil {
		return err
	}

	d.Tx = tx
	return nil
}

func (d *DatabaseClient) Commit() error {
	if d.Tx == nil {
		return errors.New("트랜잭션이 시작되지 않았습니다")
	}

	err := d.Tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseClient) CommitAndEnd() error {
	err := d.Commit()
	if err != nil {
		return err
	}
	d.EndTx()
	return nil
}

func (d *DatabaseClient) Rollback() error {
	if d.Tx == nil {
		return errors.New("트랜잭션이 시작되지 않았습니다")
	}

	err := d.Tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseClient) RollbackAndEnd() error {
	err := d.Rollback()
	if err != nil {
		return err
	}
	d.EndTx()
	return nil
}

func (d *DatabaseClient) EndTx() {
	d.Tx = nil
}

func (d *DatabaseClient) Close() {
	err := d.Client.Close()
	if err != nil {
		return
	}
}

// 마이그레이션
func (d *DatabaseClient) Migrate() {
	if err := d.Client.Schema.WriteTo(d.Ctx, os.Stdout); err != nil {
		log.Fatalf("마이그레이션 수행 중 오류 발생: %v", err)
	}
}
