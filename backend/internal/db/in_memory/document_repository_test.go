package in_memory

import (
	"context"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryDocumentRepository_GetLatestDocument(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		createdAt := time.Now().UTC()
		createdAt2 := createdAt.Add(time.Second)
		document1 := testutil.NewDocument(project.ID, []string{}, "Document 1", &createdAt)
		document2 := testutil.NewDocument(project.ID, []string{}, "Document 2", &createdAt2)
		documents := []*domain.Document{document1, document2}
		repo := NewDocumentRepository(documents)

		outputDocument, err := repo.GetLatestDocument(context.Background(), project.ID)
		assert.Nil(t, err)
		assert.Equal(t, document2, outputDocument)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		documents := []*domain.Document{}
		repo := NewDocumentRepository(documents)
		document, err := repo.GetLatestDocument(context.Background(), project.ID)
		assert.NotNil(t, err)
		assert.Nil(t, document)
	})
}
