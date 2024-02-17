package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	UserStatus   string    `json:"user_status"`
	UserType     string    `json:"user_type"`
	FailCount    int       `json:"fail_count"`
	LastLoggedIn time.Time `json:"last_logged_in"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
