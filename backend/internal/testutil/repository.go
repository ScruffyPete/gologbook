package testutil

import (
	"context"
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type FailingProjectRepo struct{}

var ErrRepoFailed = errors.New("simulated failure")

func (f *FailingProjectRepo) CreateProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	return nil, ErrRepoFailed
}
func (f *FailingProjectRepo) ListProjects(ctx context.Context) ([]*domain.Project, error) {
	return nil, ErrRepoFailed
}
func (f *FailingProjectRepo) GetProject(ctx context.Context, id string) (*domain.Project, error) {
	return nil, ErrRepoFailed
}
func (f *FailingProjectRepo) UpdateProject(ctx context.Context, id *domain.Project) error {
	return ErrRepoFailed
}
func (f *FailingProjectRepo) DeleteProject(ctx context.Context, id string) error {
	return ErrRepoFailed
}

type FailingEntryRepo struct{}

func (f *FailingEntryRepo) ListEntries(ctx context.Context, projectID string) ([]*domain.Entry, error) {
	return nil, ErrRepoFailed
}

func (f *FailingEntryRepo) CreateEntry(ctx context.Context, entry *domain.Entry) (*domain.Entry, error) {
	return nil, ErrRepoFailed
}

func (f *FailingEntryRepo) DeleteEntries(ctx context.Context, projectID string) error {
	return ErrRepoFailed
}
