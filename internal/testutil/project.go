package testutil

import (
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

func MakeDummyProjects() []domain.Project {
	projectA := domain.MakeProject("Build a treehouse")
	projectB := domain.MakeProject("Paint the garage")
	projectC := domain.MakeProject("Cook a feast")

	return []domain.Project{
		projectA,
		projectB,
		projectC,
	}
}

type FailingProjectRepo struct{}

var ErrRepoFailed = errors.New("simulated failure")

func (f *FailingProjectRepo) CreateProject(project domain.Project) error {
	return ErrRepoFailed
}
func (f *FailingProjectRepo) ListProjects() []domain.Project { return nil }
func (f *FailingProjectRepo) GetProject(id string) (*domain.Project, error) {
	return nil, ErrRepoFailed
}
func (f *FailingProjectRepo) UpdateProject(domain.Project) error {
	return ErrRepoFailed
}
func (f *FailingProjectRepo) DeleteProject(id string) error { return ErrRepoFailed }
