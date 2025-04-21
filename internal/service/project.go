package service

import (
	"errors"
	"fmt"

	"github.com/ScruffyPete/gologbook/internal/db"
	"github.com/ScruffyPete/gologbook/internal/domain"
)

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrDuplicateProject = errors.New("project not found")
)

type ProjectService struct {
	repo domain.ProjectReporitory
}

type CreateProjectInput struct {
	Title string `json:"title"`
}

func NewProjectService(repo domain.ProjectReporitory) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) ListProjects() ([]domain.Project, error) {
	projects, err := s.repo.ListProjects()
	if err != nil {
		return nil, fmt.Errorf("list projecs: %w", err)
	}
	return projects, nil
}

func (s *ProjectService) GetProject(id string) (*domain.Project, error) {
	project, err := s.repo.GetProject(id)
	if err != nil {
		return nil, fmt.Errorf("get project: %w", err)
	}
	return project, nil
}

func (s *ProjectService) CreateProject(input *CreateProjectInput) error {
	project := domain.MakeProject(input.Title)
	if err := s.repo.CreateProject(project); err != nil {
		return fmt.Errorf("create project: %w", err)
	}
	return nil
}

func (s *ProjectService) UpdateProject(id string, input *CreateProjectInput) error {
	project := domain.Project{
		ID:    id,
		Title: input.Title,
	}
	if err := s.repo.UpdateProject(project); err != nil {
		return fmt.Errorf("update project: %w", err)
	}
	return nil
}

func (s *ProjectService) DeleteProject(id string) error {
	if err := s.repo.DeleteProject(id); err != nil {
		if errors.Is(err, db.ErrProjectDoesNotExist) {
			return ErrProjectNotFound
		}
		return fmt.Errorf("get project: %w", err)
	}
	return nil
}
