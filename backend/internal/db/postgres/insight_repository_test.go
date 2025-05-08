//go:build integration

package postgres

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInsightRepository_GetInsights(t *testing.T) {
	t.Run("returns all insights for a project", func(t *testing.T) {
		project := domain.NewProject("Build a treehouse")
		insights := []*domain.Insight{
			testutil.NewInsight(project.ID, []string{}, "Insight 1"),
			testutil.NewInsight(project.ID, []string{}, "Insight 2"),
		}
		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, insights)
		defer db.Close()

		repo := NewInsightRepository(db)
		repo_insights, err := repo.ListInsights(project.ID)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(repo_insights))
		assert.Equal(t, "Insight 1", repo_insights[0].Body)
		assert.Equal(t, "Insight 2", repo_insights[1].Body)
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		db, _ := testutil.NewTestDB(nil, nil, nil, nil)
		repo := NewInsightRepository(db)
		db.Close()
		_, err := repo.ListInsights("project_id")
		assert.NotNil(t, err)
	})

	t.Run("returns an empty slice if no insights are found", func(t *testing.T) {
		project := domain.NewProject("Build a treehouse")
		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewInsightRepository(db)
		repo_insights, err := repo.ListInsights(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(repo_insights))
	})
}
