package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type EntryService struct {
	uow   domain.UnitOfWork
	queue domain.Queue
}

type CreateEntryInput struct {
	ProjectID string `json:"project_id"`
	Body      string `json:"body"`
}

func NewEntryService(uow domain.UnitOfWork, queue domain.Queue) *EntryService {
	if uow == nil {
		panic("EntryService: unit of work cannot be nil")
	}
	return &EntryService{uow: uow, queue: queue}
}

func (s *EntryService) ListEntries(ctx context.Context, projectID string) ([]*domain.Entry, error) {
	var result []*domain.Entry

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Entries.ListEntries(projectID)
		return err
	})
	if err != nil {
		slog.Error("list entries", "error", err)
		return nil, fmt.Errorf("list entries: %w", err)
	}
	return result, nil
}

func (s *EntryService) CreateEntry(
	ctx context.Context,
	input *CreateEntryInput,
) (*domain.Entry, error) {
	var result *domain.Entry

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		if _, err = repos.Projects.GetProject(input.ProjectID); err != nil {
			slog.Error("project not found", "error", err)
			return err
		}

		new_entry := domain.NewEntry(input.ProjectID, input.Body)
		result, err = repos.Entries.CreateEntry(new_entry)
		return err
	})

	if err != nil {
		slog.Error("create entry", "error", err)
		return nil, fmt.Errorf("create entry: %w", err)
	}

	if s.queue == nil {
		slog.Error("queue cannot be nil")
		return nil, fmt.Errorf("EntryService: queue cannot be nil")
	}

	key := "project_zset"
	if err := s.queue.Push(ctx, key, result.ProjectID); err != nil {
		slog.Error("push message to queue", "error", err)
		return nil, fmt.Errorf("push message to queue: %w", err)
	}
	return result, nil
}
