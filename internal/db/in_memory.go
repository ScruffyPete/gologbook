package db

import (
	"errors"
)

type inMemoryProjectRepository struct {
	projects map[string]Project
}

var ErrProjectExists = errors.New("project already exists")
var ErrProjectDoesNotExist = errors.New("proejct doesn't exist")

func newInMemoryProjectRepository(data map[string]Project) *inMemoryProjectRepository {
	if data == nil {
		data = make(map[string]Project)
	}

	return &inMemoryProjectRepository{projects: data}
}

func (repo *inMemoryProjectRepository) ListProjects() []Project {
	projects := make([]Project, 0, len(repo.projects))

	for _, project := range repo.projects {
		projects = append(projects, project)
	}

	return projects
}

func (repo *inMemoryProjectRepository) GetProject(id string) *Project {
	if projectData, exists := repo.projects[id]; exists {
		return &projectData
	}

	return nil
}

func (repo *inMemoryProjectRepository) CreateProject(project Project) error {
	if _, exists := repo.projects[project.ID]; exists {
		return ErrProjectExists
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *inMemoryProjectRepository) UpdateProject(project Project) error {
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
