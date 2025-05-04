package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUser(email string, password string) *User {
	return &User{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.999999Z"),
		Email:     email,
		Password:  password,
	}
}
