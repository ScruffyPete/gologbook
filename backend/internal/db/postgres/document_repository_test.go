//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDocumentRepository_GetDocument(t *testing.T) {
	t.Run("returns latest documents for a project", func(t *testing.T) {
		project := domain.NewProject("Build a treehouse")
		createdAt := time.Now().UTC()
		createdAt2 := createdAt.Add(time.Second)
		document1 := testutil.NewDocument(project.ID, []string{}, "Document 1", &createdAt)
		document2 := testutil.NewDocument(project.ID, []string{}, "Document 2", &createdAt2)
		documents := []*domain.Document{document1, document2}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, nil, documents)

		repo := NewDocumentRepository(db)
		outputDocument, err := repo.GetLatestDocument(ctx, project.ID)

		assert.Nil(t, err)
		assert.Equal(t, document2, outputDocument)
	})

	t.Run("returns an empty slice if no documents are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Build a treehouse")
		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, nil, nil)

		repo := NewDocumentRepository(db)
		outputDocument, err := repo.GetLatestDocument(ctx, project.ID)
		assert.NotNil(t, err)
		assert.Nil(t, outputDocument)
	})
}
