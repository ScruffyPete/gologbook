package in_memory

import (
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryInsightRepository struct {
	insights map[string]*domain.Insight
}

func NewInsightRepository(insights []*domain.Insight) *InMemoryInsightRepository {
	data := make(map[string]*domain.Insight)

	for _, insight := range insights {
		data[insight.ID] = insight
	}

	return &InMemoryInsightRepository{insights: data}
}

func (repo *InMemoryInsightRepository) ListInsights(projectID string) ([]*domain.Insight, error) {
	insights := make([]*domain.Insight, 0)
	for _, insight := range repo.insights {
		if insight.ProjectID == projectID {
			insights = append(insights, insight)
		}
	}
	sort.Slice(insights, func(i, j int) bool {
		return insights[j].CreatedAt > insights[i].CreatedAt
	})
	return insights, nil
}
