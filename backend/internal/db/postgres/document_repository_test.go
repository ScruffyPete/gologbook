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
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	project := domain.NewProject("Build a treehouse")
	createdAt := time.Now().UTC()
	createdAt2 := createdAt.Add(time.Second)
	document1 := testutil.NewDocument(project.ID, []string{}, "Document 1", &createdAt)
	document2 := testutil.NewDocument(project.ID, []string{}, "Document 2", &createdAt2)
	documents := []*domain.Document{document1, document2}

	t.Run("returns latest documents for a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var outputDocument *domain.Document
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}
			for _, d := range documents {
				if _, err := repos.Documents.CreateDocument(ctx, d); err != nil {
					return err
				}
			}

			outputDocument, err = repos.Documents.GetLatestDocument(ctx, project.ID)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, document2, outputDocument)
	})

	t.Run("returns an error if no documents are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}

			_, err = repos.Documents.GetLatestDocument(ctx, project.ID)
			return err
		})
		assert.NotNil(t, err)
	})
}
