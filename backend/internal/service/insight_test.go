package service

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestInsightService_ListInsights(t *testing.T) {
	project := domain.NewProject("Test Project")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	entry := domain.NewEntry("Test Entry", project.ID)
	entryRepo := in_memory.NewEntryRepository([]*domain.Entry{entry})

	ctx := context.Background()

	t.Run("should return insights", func(t *testing.T) {
		body := "Test Insight"
		insight := testutil.NewInsight(project.ID, []string{entry.ID}, body, nil)
		insightRepo := in_memory.NewInsightRepository([]*domain.Insight{insight})
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
			Insights: insightRepo,
		}
		svc := NewInsightService(&uow)

		insights, err := svc.ListInsights(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, []*domain.Insight{insight}, insights)
	})

	t.Run("should return empty list if no insights", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		svc := NewInsightService(uow)

		insights, err := svc.ListInsights(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, []*domain.Insight{}, insights)
	})
}
