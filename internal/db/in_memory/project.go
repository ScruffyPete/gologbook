package in_memory

import (
	"errors"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type projectRepository struct {
	projects map[string]domain.Project
}

var (
	ErrDuplicateProject    = errors.New("project already exists")
	ErrProjectDoesNotExist = errors.New("project doesn't exist")
)

func NewProjectRepository(projects []domain.Project) *projectRepository {
	data := make(map[string]domain.Project)

	for _, p := range projects {
		data[p.ID] = domain.Project{
			ID:    p.ID,
			Title: p.Title,
		}
	}

	return &projectRepository{projects: data}
}

func (repo *projectRepository) ListProjects() ([]domain.Project, error) {
	projects := make([]domain.Project, 0, len(repo.projects))

	for _, project := range repo.projects {
		projects = append(projects, project)
	}

	return projects, nil
}

func (repo *projectRepository) GetProject(id string) (*domain.Project, error) {
	if projectData, exists := repo.projects[id]; exists {
		return &projectData, nil
	}

	return nil, ErrProjectDoesNotExist
}

func (repo *projectRepository) CreateProject(project domain.Project) error {
	if _, exists := repo.projects[project.ID]; exists {
		return ErrDuplicateProject
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *projectRepository) UpdateProject(project domain.Project) error {
	if _, exists := repo.projects[project.ID]; !exists {
		return ErrProjectDoesNotExist
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *projectRepository) DeleteProject(id string) error {
	if _, exists := repo.projects[id]; !exists {
		return ErrProjectDoesNotExist
	}
	delete(repo.projects, id)
	return nil
}
