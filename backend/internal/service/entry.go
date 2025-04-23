package service

import (
	"context"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type EntryService struct {
	uow domain.UnitOfWork
}

type CreateEntryInput struct {
	Body string `json:"body"`
}

func NewEntryService(uow domain.UnitOfWork) *EntryService {
	if uow == nil {
		panic("EntryService: unit of work cannot be nil")
	}
	return &EntryService{uow: uow}
}

func (s *EntryService) ListEntries(ctx context.Context, projectID string) ([]*domain.Entry, error) {
	var result []*domain.Entry

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Entries.ListEntries(projectID)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("list entries: %w", err)
	}
	return result, nil
}

func (s *EntryService) CreateEntry(
	ctx context.Context,
	projectID string,
	input *CreateEntryInput,
) (*domain.Entry, error) {
	var result *domain.Entry

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		if _, err = repos.Projects.GetProject(projectID); err != nil {
			return err
		}

		new_entry := domain.MakeEntry(projectID, input.Body)
		result, err = repos.Entries.CreateEntry(new_entry)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("create entry: %w", err)
	}

	return result, nil
}
