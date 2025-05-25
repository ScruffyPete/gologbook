package service

import (
	"context"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type DocumentService struct {
	uow domain.UnitOfWork
}

func NewDocumentService(uow domain.UnitOfWork) *DocumentService {
	return &DocumentService{uow: uow}
}

func (s *DocumentService) ListDocuments(ctx context.Context, projectID string) ([]*domain.Document, error) {
	var result []*domain.Document

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Documents.ListDocuments(projectID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("list docuemnts: %w", err)
	}

	return result, nil
}
