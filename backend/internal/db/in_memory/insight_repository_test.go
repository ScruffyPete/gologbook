package in_memory

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryInsightRepository_GetInsights(t *testing.T) {
	t.Run("project wide insights", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		insights := []*domain.Insight{
			testutil.NewInsight(project.ID, []string{}, "Insight 1"),
			testutil.NewInsight(project.ID, []string{}, "Insight 2"),
		}
		repo := NewInsightRepository(insights)
		insights, err := repo.ListInsights(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(insights), 2)
		assert.Equal(t, insights[0].Body, "Insight 1")
		assert.Equal(t, insights[1].Body, "Insight 2")
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		insights := []*domain.Insight{}
		repo := NewInsightRepository(insights)
		insights, err := repo.ListInsights(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(insights), 0)
	})
}
