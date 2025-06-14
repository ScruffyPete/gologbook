package in_memory

import (
	"context"
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type projectRepository struct {
	projects map[string]*domain.Project
}

func NewProjectRepository(projects []*domain.Project) *projectRepository {
	data := make(map[string]*domain.Project)

	for _, p := range projects {
		data[p.ID] = p
	}

	return &projectRepository{projects: data}
}

func (repo *projectRepository) ListProjects(ctx context.Context) ([]*domain.Project, error) {
	projects := make([]*domain.Project, 0, len(repo.projects))

	for _, p := range repo.projects {
		projects = append(projects, p)
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].CreatedAt > projects[j].CreatedAt
	})

	return projects, nil
}

func (repo *projectRepository) GetProject(ctx context.Context, id string) (*domain.Project, error) {
	if projectData, exists := repo.projects[id]; exists {
		return projectData, nil
	}

	return nil, domain.NewErrProjectDoesNotExist(id)
}

func (repo *projectRepository) CreateProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	repo.projects[project.ID] = project
	return project, nil
}

func (repo *projectRepository) UpdateProject(ctx context.Context, project *domain.Project) error {
	if _, exists := repo.projects[project.ID]; !exists {
		return domain.NewErrProjectDoesNotExist(project.ID)
	}
	repo.projects[project.ID] = project
	return nil
}

func (repo *projectRepository) DeleteProject(ctx context.Context, id string) error {
	if _, exists := repo.projects[id]; !exists {
		return domain.NewErrProjectDoesNotExist(id)
	}
	delete(repo.projects, id)
	return nil
}
