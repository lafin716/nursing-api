package database

import (
	"fmt"
	"log"
	"nursing_api/ent"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func (c *Config) Serialize() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		c.Host,
		c.Port,
		c.User,
		c.Database,
		c.Password)
}

func NewPostgresClient(config *Config) (*ent.Client, error) {
	client, err := ent.Open("postgres", config.Serialize())
	if err != nil {
		log.Printf("데이터베이스 연결에 실패하였습니다: %v", err)
		return nil, err
	}

	return client, nil
}
