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

func (s *EntryService) CreateEntry(projectID string, input *CreateEntryInput) (*domain.Entry, error) {
	_, err := s.projectRepo.GetProject(projectID)
	if err != nil {
		return nil, fmt.Errorf("crete entry: %w", err)
	}

	new_entry := domain.MakeEntry(projectID, input.Body)
	entry, err := s.entryRepo.CreateEntry(new_entry)
	if err != nil {
		return nil, fmt.Errorf("crete entry: %w", err)
	}

	return entry, nil
}
