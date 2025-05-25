package service

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDocumentService_ListDocuments(t *testing.T) {
	project := domain.NewProject("Test Project")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	entry := domain.NewEntry("Test Entry", project.ID)
	entryRepo := in_memory.NewEntryRepository([]*domain.Entry{entry})

	ctx := context.Background()

	t.Run("should return documents", func(t *testing.T) {
		body := "Test Document"
		document := testutil.NewDocument(project.ID, []string{entry.ID}, body, nil)
		documentRepo := in_memory.NewDocumentRepository([]*domain.Document{document})
		uow := in_memory.InMemoryUnitOfWork{
			Projects:  projectRepo,
			Entries:   entryRepo,
			Documents: documentRepo,
		}
		svc := NewDocumentService(&uow)

		documents, err := svc.ListDocuments(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, []*domain.Document{document}, documents)
	})

	t.Run("should return empty list if no docuements", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		svc := NewDocumentService(uow)

		documents, err := svc.ListDocuments(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, []*domain.Document{}, documents)
	})
}
