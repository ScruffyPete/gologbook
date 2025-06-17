package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (repo *DocumentRepository) GetLatestDocument(ctx context.Context, projectID string) (*domain.Document, error) {
	row := repo.db.QueryRowContext(
		ctx,
		`SELECT id, created_at, project_id, entry_ids, body 
		 FROM documents 
		 WHERE project_id = $1 
		 ORDER BY created_at DESC 
		 LIMIT 1`,
		projectID,
	)

	var document domain.Document
	var entryIDsRaw []byte
	if err := row.Scan(&document.ID, &document.CreatedAt, &document.ProjectID, &entryIDsRaw, &document.Body); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no document found")
		}
		return nil, err
	}
	if err := json.Unmarshal(entryIDsRaw, &document.EntryIDs); err != nil {
		return nil, err
	}

	return &document, nil
}
