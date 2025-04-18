package db

import "time"

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
