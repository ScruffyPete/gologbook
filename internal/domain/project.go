package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string
	CreatedAt time.Time
	Title     string
}

func MakeProject(title string) *Project {
	return &Project{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Title:     title,
	}
}
