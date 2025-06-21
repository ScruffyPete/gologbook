package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

type DocumentRepository struct {
	tx *sql.Tx
}

func NewDocumentRepository(tx *sql.Tx) *DocumentRepository {
	return &DocumentRepository{tx: tx}
}

func (repo *DocumentRepository) CreateDocument(ctx context.Context, document *domain.Document) (*domain.Document, error) {
	// For test purposes
	entryIDsJSON, err := json.Marshal(document.EntryIDs)
	if err != nil {
		return nil, err
	}
	_, err = repo.tx.ExecContext(
		ctx,
		"INSERT INTO documents (id, created_at, project_id, entry_ids, body) VALUES ($1, $2, $3, $4, $5)",
		document.ID,
		document.CreatedAt,
		document.ProjectID,
		entryIDsJSON,
		document.Body,
	)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (repo *DocumentRepository) GetLatestDocument(ctx context.Context, projectID string) (*domain.Document, error) {
	row := repo.tx.QueryRowContext(
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
			return nil, fmt.Errorf("no documents found for projectID: %s", projectID)
		}
		return nil, err
	}
	if err := json.Unmarshal(entryIDsRaw, &document.EntryIDs); err != nil {
		return nil, err
	}

	return &document, nil
}
