package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
}

func NewProject(title string) *Project {
	return &Project{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.999999Z"),
		Title:     title,
	}
}

type Entry struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	ProjectID string `json:"project_id"`
	Body      string `json:"body"`
}

func NewEntry(projectID string, body string) *Entry {
	return &Entry{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC().Format("2006-01-02T15:04:05.999999Z"),
		ProjectID: projectID,
		Body:      body,
	}
}

type Insight struct {
	ID        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	ProjectID string   `json:"project_id"`
	EntryIDs  []string `json:"entry_ids,omitempty"`
	Body      string   `json:"body"`
}
