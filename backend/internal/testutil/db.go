package testutil

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

func NewTestDB(
	t testing.TB,
	ctx context.Context,
	users []*domain.User,
	projects []*domain.Project,
	entries []*domain.Entry,
	documents []*domain.Document,
) *sql.DB {
	t.Helper()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	// Ensure clean state before inserting test data
	if _, err := db.ExecContext(ctx, "TRUNCATE TABLE users, projects, entries, documents RESTART IDENTITY CASCADE;"); err != nil {
		t.Fatalf("failed to clean up database: %v", err)
	}

	for _, user := range users {
		_, err := db.ExecContext(
			ctx,
			"INSERT INTO users (id, created_at, email, password) VALUES ($1, $2, $3, $4)",
			user.ID,
			user.CreatedAt,
			user.Email,
			user.Password,
		)
		if err != nil {
			t.Fatalf("failed to insert user: %v", err)
		}
	}

	for _, project := range projects {
		_, err := db.ExecContext(
			ctx,
			"INSERT INTO projects (id, title, created_at) VALUES ($1, $2, $3)",
			project.ID,
			project.Title,
			project.CreatedAt,
		)
		if err != nil {
			t.Fatalf("failed to insert project: %v", err)
		}
	}

	for _, entry := range entries {
		_, err := db.ExecContext(
			ctx,
			"INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)",
			entry.ID,
			entry.CreatedAt,
			entry.ProjectID,
			entry.Body,
		)
		if err != nil {
			t.Fatalf("failed to insert entry: %v", err)
		}
	}

	for _, document := range documents {
		entryIDsJSON, err := json.Marshal(document.EntryIDs)
		if err != nil {
			db.Close()
			t.Fatalf("failed to marshal entry IDs: %v", err)
		}
		_, err = db.ExecContext(
			ctx,
			"INSERT INTO documents (id, created_at, project_id, entry_ids, body) VALUES ($1, $2, $3, $4, $5)",
			document.ID,
			document.CreatedAt,
			document.ProjectID,
			entryIDsJSON,
			document.Body,
		)
		if err != nil {
			t.Fatalf("failed to insert document: %v", err)
		}
	}

	return db
}
