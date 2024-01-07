package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name         string    `db:"name" json:"name" validate:"required,lte=100"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password_hash" json:"-" validate:"required,lte=255"`
	UserStatus   string    `db:"user_status" json:"user_status" validate:"required"`
	UserType     string    `db:"user_type" json:"user_type"`
	FailCount    int       `db:"fail_count" json:"fail_count"`
	LastLoggedIn time.Time `db:"last_logged_in" json:"last_logged_in"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
