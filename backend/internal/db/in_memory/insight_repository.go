package in_memory

import (
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryInsightRepository struct {
	insights map[string][]*domain.Insight
}

func NewInsightRepository(insights []*domain.Insight) *InMemoryInsightRepository {
	data := make(map[string][]*domain.Insight)

	for _, insight := range insights {
		data[insight.ProjectID] = append(data[insight.ProjectID], insight)
	}

	return &InMemoryInsightRepository{insights: data}
}

func (repo *InMemoryInsightRepository) ListInsights(projectID string) ([]*domain.Insight, error) {
	insights := repo.insights[projectID]

	sorted := make([]*domain.Insight, len(insights))
	copy(sorted, insights)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CreatedAt < sorted[j].CreatedAt
	})

	return sorted, nil
}
