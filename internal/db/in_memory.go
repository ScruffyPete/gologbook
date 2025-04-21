package db

import (
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type inMemoryProjectRepository struct {
	projects map[string]domain.Project
}

var (
	ErrDuplicateProject    = errors.New("project already exists")
	ErrProjectDoesNotExist = errors.New("project doesn't exist")
)

func NewInMemoryProjectRepository(projects []domain.Project) *inMemoryProjectRepository {
	data := make(map[string]domain.Project)

	for _, p := range projects {
		data[p.ID] = domain.Project{
			ID:    p.ID,
			Title: p.Title,
		}
	}

	return &inMemoryProjectRepository{projects: data}
}

func (repo *inMemoryProjectRepository) ListProjects() []domain.Project {
	projects := make([]domain.Project, 0, len(repo.projects))

	for _, project := range repo.projects {
		projects = append(projects, project)
	}

	return projects
}

func (repo *inMemoryProjectRepository) GetProject(id string) (*domain.Project, error) {
	if projectData, exists := repo.projects[id]; exists {
		return &projectData, nil
	}

	return nil, ErrProjectDoesNotExist
}

func (repo *inMemoryProjectRepository) CreateProject(project domain.Project) error {
	if _, exists := repo.projects[project.ID]; exists {
		return ErrDuplicateProject
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *inMemoryProjectRepository) UpdateProject(project domain.Project) error {
	if _, exists := repo.projects[project.ID]; !exists {
		return ErrProjectDoesNotExist
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *inMemoryProjectRepository) DeleteProject(id string) error {
	if _, exists := repo.projects[id]; !exists {
		return ErrProjectDoesNotExist
	}
	delete(repo.projects, id)
	return nil
}
