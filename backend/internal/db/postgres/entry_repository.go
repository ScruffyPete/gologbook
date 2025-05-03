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
	return nil, nil
}

func (repo *entryRepository) CreateEntry(entry *domain.Entry) (*domain.Entry, error) {
	return nil, nil
}

func (repo *entryRepository) DeleteEntries(projectID string) error {
	return nil
}
