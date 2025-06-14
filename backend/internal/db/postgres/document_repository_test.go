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
	t.Run("returns all documents for a project", func(t *testing.T) {
		project := domain.NewProject("Build a treehouse")
		createdAt := time.Now().UTC()
		createdAt2 := createdAt.Add(time.Second)
		documents := []*domain.Document{
			testutil.NewDocument(project.ID, []string{}, "Document 1", &createdAt),
			testutil.NewDocument(project.ID, []string{}, "Document 2", &createdAt2),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, documents)
		defer db.Close()

		repo := NewDocumentRepository(db)
		repo_documents, err := repo.ListDocuments(project.ID)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(repo_documents))
		assert.Equal(t, "Document 1", repo_documents[0].Body)
		assert.Equal(t, "Document 2", repo_documents[1].Body)
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		repo := NewDocumentRepository(db)
		db.Close()
		_, err := repo.ListDocuments("project_id")
		assert.NotNil(t, err)
	})

	t.Run("returns an empty slice if no documents are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Build a treehouse")
		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewDocumentRepository(db)
		repo_documents, err := repo.ListDocuments(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(repo_documents))
	})
}
