package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID    string
	Title string
}

type Entry struct {
	ID        string
	ProjectID string
	Author    string // TODO proabably needs a user reference
	Body      string
	CratedAt  time.Time
}

func MakeProject(title string) Project {
	return Project{
		ID:    uuid.NewString(),
		Title: title,
	}
}
