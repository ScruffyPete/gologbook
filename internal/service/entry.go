package service

import (
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type EntryService struct {
	entryRepo   domain.EntryRepository
	projectRepo domain.ProjectReporitory
}

type CreateEntryInput struct {
	Body string `json:"body"`
}

func NewEntryService(
	entryRepo domain.EntryRepository,
	projectRepo domain.ProjectReporitory,
) *EntryService {
	return &EntryService{entryRepo, projectRepo}
}

func (s *EntryService) ListEntries(projectID string) ([]*domain.Entry, error) {
	entries, err := s.entryRepo.ListEntries(projectID)
	if err != nil {
		return nil, fmt.Errorf("list entries: %w", err)
	}
	return entries, nil
}

func (s *EntryService) CreateEntry(projectID string, input *CreateEntryInput) error {
	_, err := s.projectRepo.GetProject(projectID)
	if err != nil {
		return fmt.Errorf("crete entry: %w", err)
	}

	entry := domain.MakeEntry(projectID, input.Body)
	if err := s.entryRepo.CreateEntry(entry); err != nil {
		return fmt.Errorf("crete entry: %w", err)
	}

	return nil
}
