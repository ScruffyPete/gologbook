package in_memory

import (
	"context"
	"errors"

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

func (repo *InMemoryDocumentRepository) GetLatestDocument(ctx context.Context, projectID string) (*domain.Document, error) {
	documents := repo.documents[projectID]

	if len(documents) == 0 {
		return nil, errors.New("no document found")
	}

	latest := documents[0]
	for _, doc := range documents[1:] {
		if doc.CreatedAt > latest.CreatedAt {
			latest = doc
		}
	}

	return latest, nil
}
