package service

import (
	"context"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type DocumentService struct {
	uow   domain.UnitOfWork
	queue domain.Queue
}

func NewDocumentService(uow domain.UnitOfWork, queue domain.Queue) *DocumentService {
	return &DocumentService{uow: uow, queue: queue}
}

func (s *DocumentService) GetLatestDocument(ctx context.Context, projectID string) (*domain.Document, error) {
	var result *domain.Document

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		if _, err = repos.Projects.GetProject(ctx, projectID); err != nil {
			return err
		}
		result, err = repos.Documents.GetLatestDocument(ctx, projectID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("list docuemnts: %w", err)
	}

	return result, nil
}

func (s *DocumentService) ConsumeDocumentStream(ctx context.Context, projectID string) <-chan string {
	return s.queue.SubscribeForDocumentTokens(ctx, projectID)
}
