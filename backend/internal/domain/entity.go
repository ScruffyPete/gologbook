package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
}

func MakeProject(title string) *Project {
	return &Project{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		Title:     title,
	}
}

type Entry struct {
	ID        string    `json:"id"`
	CratedAt  time.Time `json:"createdAt"`
	ProjectID string    `json:"projectId"`
	Body      string    `json:"body"`
}

func MakeEntry(projectID string, body string) *Entry {
	return &Entry{
		ID:        uuid.NewString(),
		CratedAt:  time.Now(),
		ProjectID: projectID,
		Body:      body,
	}
}
