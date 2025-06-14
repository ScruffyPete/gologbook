package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrDuplicateProject = errors.New("project not found")
)

type ProjectService struct {
	uow domain.UnitOfWork
}

type CreateProjectInput struct {
	Title string `json:"title"`
}

func NewProjectService(uow domain.UnitOfWork) *ProjectService {
	if uow == nil {
		panic("ProjectService: unit of work cannot be nil")
	}
	return &ProjectService{uow: uow}
}

func (s *ProjectService) ListProjects(ctx context.Context) ([]*domain.Project, error) {
	var result []*domain.Project

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Projects.ListProjects(ctx)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("list projecs: %w", err)
	}
	return result, nil
}

func (s *ProjectService) GetProject(ctx context.Context, id string) (*domain.Project, error) {
	var result *domain.Project

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		result, err = repos.Projects.GetProject(ctx, id)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("get project: %w", err)
	}
	return result, nil
}

func (s *ProjectService) CreateProject(ctx context.Context, input *CreateProjectInput) (*domain.Project, error) {
	var result *domain.Project

	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		var err error
		new_project := domain.NewProject(input.Title)
		result, err = repos.Projects.CreateProject(ctx, new_project)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}

	return result, nil
}

func (s *ProjectService) UpdateProject(
	ctx context.Context,
	id string,
	input *CreateProjectInput,
) error {
	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		if project, err := repos.Projects.GetProject(ctx, id); err != nil {
			return err
		} else {
			project.Title = input.Title
			return repos.Projects.UpdateProject(ctx, project)
		}

	})

	if err != nil {
		return fmt.Errorf("update project: %w", err)
	}

	return nil
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	err := s.uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		return repos.Projects.DeleteProject(ctx, id)
	})

	if err != nil {
		return fmt.Errorf("delete project: %w", err)
	}
	return nil
}
