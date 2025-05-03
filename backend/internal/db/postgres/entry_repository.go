package postgres

import (
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type entryRepository struct {
	db *sql.DB
}

func NewEntryRepository(db *sql.DB) *entryRepository {
	return &entryRepository{db: db}
}

func (repo *entryRepository) ListEntries(projectID string) ([]*domain.Entry, error) {
	rows, err := repo.db.Query(
		"SELECT id, created_at, project_id, body FROM entries WHERE project_id = $1 ORDER BY created_at DESC",
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

func (repo *entryRepository) CreateEntry(entry *domain.Entry) (*domain.Entry, error) {
	_, err := repo.db.Exec(
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

func (repo *entryRepository) DeleteEntries(projectID string) error {
	_, err := repo.db.Exec(
		"DELETE FROM entries WHERE project_id = $1",
		projectID,
	)
	if err != nil {
		return err
	}

	return nil
}
