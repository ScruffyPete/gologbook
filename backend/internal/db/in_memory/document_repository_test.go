package in_memory

import (
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryDocumentRepository_ListDocuments(t *testing.T) {
	t.Run("project wide documents", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		createdAt := time.Now().UTC()
		createdAt2 := createdAt.Add(time.Second)
		documents := []*domain.Document{
			testutil.NewDocument(project.ID, []string{}, "Document 1", &createdAt),
			testutil.NewDocument(project.ID, []string{}, "Document 2", &createdAt2),
		}
		repo := NewDocumentRepository(documents)
		documents, err := repo.ListDocuments(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(documents), 2)
		assert.Equal(t, documents[0].Body, "Document 1")
		assert.Equal(t, documents[1].Body, "Document 2")
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		documents := []*domain.Document{}
		repo := NewDocumentRepository(documents)
		documents, err := repo.ListDocuments(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(documents), 0)
	})
}
