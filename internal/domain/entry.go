package domain

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID        string
	CratedAt  time.Time
	ProjectID string
	// Author    string // TODO proabably needs a user reference
	Body string
}

func MakeEntry(projectID string, body string) Entry {
	return Entry{
		ID:        uuid.NewString(),
		CratedAt:  time.Now(),
		ProjectID: projectID,
		Body:      body,
	}
}
