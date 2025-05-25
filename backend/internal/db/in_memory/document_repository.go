package in_memory

import (
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryDocumentRepository struct {
	documents map[string][]*domain.Document
}

func NewDocumentRepository(documents []*domain.Document) *InMemoryDocumentRepository {
	data := make(map[string][]*domain.Document)

	for _, document := range documents {
		data[document.ProjectID] = append(data[document.ProjectID], document)
	}

	return &InMemoryDocumentRepository{documents: data}
}

func (repo *InMemoryDocumentRepository) ListDocuments(projectID string) ([]*domain.Document, error) {
	documents := repo.documents[projectID]

	sorted := make([]*domain.Document, len(documents))
	copy(sorted, documents)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CreatedAt < sorted[j].CreatedAt
	})

	return sorted, nil
}
