package in_memory

import "github.com/ScruffyPete/gologbook/internal/domain"

type entryRepository struct {
	entries map[string]*domain.Entry
}

func NewEntryRepository(entries []*domain.Entry) *entryRepository {
	data := make(map[string]*domain.Entry)

	for _, e := range entries {
		data[e.ID] = e
	}

	return &entryRepository{entries: data}
}

func (repo *entryRepository) ListEntries() ([]*domain.Entry, error) {
	entries := make([]*domain.Entry, 0, len(repo.entries))

	for _, e := range repo.entries {
		entries = append(entries, e)
	}

	return entries, nil
}

func (repo *entryRepository) CreateEntry(entry *domain.Entry) error {
	repo.entries[entry.ID] = entry
	return nil
}
