package service

import (
	"context"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InsightService struct {
	uow domain.UnitOfWork
}

func NewInsightService(uow domain.UnitOfWork) *InsightService {
	return &InsightService{uow: uow}
}

func (s *InsightService) ListInsights(ctx context.Context, projectID string) ([]*domain.Insight, error) {
	var result []*domain.Insight

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Insights.ListInsights(projectID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("list insights: %w", err)
	}

	return result, nil
}
