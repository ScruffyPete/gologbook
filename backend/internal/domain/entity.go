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

type Entry struct {
	ID        string
	CratedAt  time.Time
	ProjectID string
	Body      string
}

func MakeEntry(projectID string, body string) *Entry {
	return &Entry{
		ID:        uuid.NewString(),
		CratedAt:  time.Now(),
		ProjectID: projectID,
		Body:      body,
	}
}
