package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (repo *DocumentRepository) ListDocuments(projectID string) ([]*domain.Document, error) {
	rows, err := repo.db.Query(
		"SELECT id, created_at, project_id, entry_ids, body FROM documents WHERE project_id = $1",
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	documents := make([]*domain.Document, 0)
	for rows.Next() {
		var document domain.Document
		var entryIDsRaw []byte
		if err := rows.Scan(&document.ID, &document.CreatedAt, &document.ProjectID, &entryIDsRaw, &document.Body); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(entryIDsRaw, &document.EntryIDs); err != nil {
			return nil, err
		}
		documents = append(documents, &document)
	}
	return documents, nil
}
