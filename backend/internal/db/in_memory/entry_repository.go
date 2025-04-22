package in_memory

import (
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type entryRepository struct {
	entries map[string][]*domain.Entry
}

func NewEntryRepository(entries []*domain.Entry) *entryRepository {
	data := make(map[string][]*domain.Entry)

	for _, e := range entries {
		data[e.ProjectID] = append(data[e.ProjectID], e)
	}

	return &entryRepository{entries: data}
}

func (repo *entryRepository) ListEntries(projectID string) ([]*domain.Entry, error) {
	entries := repo.entries[projectID]

	sorted := make([]*domain.Entry, len(entries))
	copy(sorted, entries)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CratedAt.Before(sorted[j].CratedAt)
	})

	return sorted, nil
}

func (repo *entryRepository) CreateEntry(entry *domain.Entry) error {
	entries := repo.entries[entry.ProjectID]
	entries = append(entries, entry)
	repo.entries[entry.ProjectID] = entries
	return nil
}
