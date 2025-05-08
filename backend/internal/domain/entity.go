package domain

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
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
	CreatedAt string `json:"createdAt"`
	ProjectID string `json:"projectId"`
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
	CreatedAt string   `json:"createdAt"`
	ProjectID string   `json:"projectId"`
	EntryIDs  []string `json:"entryIds,omitempty"`
	Body      string   `json:"body"`
}
