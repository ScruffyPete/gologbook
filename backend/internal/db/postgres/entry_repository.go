package postgres

import (
	"context"
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type entryRepository struct {
	tx *sql.Tx
}

func NewEntryRepository(tx *sql.Tx) *entryRepository {
	return &entryRepository{tx: tx}
}

func (repo *entryRepository) ListEntries(ctx context.Context, projectID string) ([]*domain.Entry, error) {
	rows, err := repo.tx.QueryContext(
		ctx,
		"SELECT id, created_at, project_id, body FROM entries WHERE project_id = $1 ORDER BY created_at ASC",
		projectID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]*domain.Entry, 0)
	for rows.Next() {
		var entry domain.Entry
		if err := rows.Scan(&entry.ID, &entry.CreatedAt, &entry.ProjectID, &entry.Body); err != nil {
			return nil, err
		}
		entries = append(entries, &entry)
	}

	return entries, nil
}

func (repo *entryRepository) CreateEntry(ctx context.Context, entry *domain.Entry) (*domain.Entry, error) {
	_, err := repo.tx.ExecContext(
		ctx,
		"INSERT INTO entries (id, created_at, project_id, body) VALUES ($1, $2, $3, $4)",
		entry.ID,
		entry.CreatedAt,
		entry.ProjectID,
		entry.Body,
	)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (repo *entryRepository) DeleteEntries(ctx context.Context, projectID string) error {
	_, err := repo.tx.ExecContext(
		ctx,
		"DELETE FROM entries WHERE project_id = $1",
		projectID,
	)
	if err != nil {
		return err
	}

	return nil
}
