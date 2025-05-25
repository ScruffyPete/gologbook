package testutil

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

func NewTestDB(
	users []*domain.User,
	projects []*domain.Project,
	entries []*domain.Entry,
	documents []*domain.Document,
) (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Ensure clean state before inserting test data
	if _, err := db.Exec("TRUNCATE TABLE users, projects, entries, documents RESTART IDENTITY CASCADE;"); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to clean up database: %w", err)
	}

	for _, user := range users {
		_, err := db.Exec(
			"INSERT INTO users (id, created_at, email, password) VALUES ($1, $2, $3, $4)",
			user.ID,
			user.CreatedAt,
			user.Email,
			user.Password,
		)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to insert user: %w", err)
		}
	}

	for _, project := range projects {
		_, err := db.Exec(
			"INSERT INTO projects (id, title, created_at) VALUES ($1, $2, $3)",
			project.ID,
			project.Title,
			project.CreatedAt,
		)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to insert project: %w", err)
		}
	}

	for _, entry := range entries {
		_, err := db.Exec(
			"INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)",
			entry.ID,
			entry.CreatedAt,
			entry.ProjectID,
			entry.Body,
		)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to insert entry: %w", err)
		}
	}

	for _, document := range documents {
		entryIDsJSON, err := json.Marshal(document.EntryIDs)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to marshal entry IDs: %w", err)
		}
		_, err = db.Exec(
			"INSERT INTO documents (id, created_at, project_id, entry_ids, body) VALUES ($1, $2, $3, $4, $5)",
			document.ID,
			document.CreatedAt,
			document.ProjectID,
			entryIDsJSON,
			document.Body,
		)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to insert document: %w", err)
		}
	}

	return db, nil
}
