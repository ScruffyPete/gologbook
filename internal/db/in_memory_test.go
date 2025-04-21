package db

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListProjects(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		projects := testutil.MakeDummyProjects()
		repo := NewInMemoryProjectRepository(projects)

		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_projects, projects)
	})

	t.Run("empty data", func(t *testing.T) {
		projects := []domain.Project{}
		repo := NewInMemoryProjectRepository(projects)
		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_projects, projects)
	})
}

func TestGetProject(t *testing.T) {

	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		repo := NewInMemoryProjectRepository([]domain.Project{project})
		repo_project, err := repo.GetProject(project.ID)
		assert.Equal(t, repo_project, &project)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("invalid project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository(testutil.MakeDummyProjects())
		repo_project, err := repo.GetProject(uuid.NewString())
		assert.Nil(t, repo_project)
		assert.ErrorIs(t, err, ErrProjectDoesNotExist)
	})
}

func TestCreateProject(t *testing.T) {
	project := domain.MakeProject("Write a novel")

	t.Run("new project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository(nil)
		err := repo.CreateProject(project)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("existing project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository([]domain.Project{project})
		err := repo.CreateProject(project)
		assert.ErrorIs(t, err, ErrDuplicateProject)
	})
}

func TestUpdateProject(t *testing.T) {
	project := domain.MakeProject("Throw a ball")

	t.Run("existing project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository([]domain.Project{project})
		err := repo.UpdateProject(domain.Project{ID: project.ID, Title: "Throw THE ball"})
		assert.ErrorIs(t, err, nil)
	})

	t.Run("missing project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository(testutil.MakeDummyProjects())
		err := repo.UpdateProject(project)
		assert.ErrorIs(t, err, ErrProjectDoesNotExist)
	})
}

func TestDeleteProject(t *testing.T) {
	project := domain.MakeProject("Earn a million")

	t.Run("existing project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository([]domain.Project{project})
		err := repo.DeleteProject(project.ID)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("invalid project", func(t *testing.T) {
		repo := NewInMemoryProjectRepository(testutil.MakeDummyProjects())
		err := repo.DeleteProject(project.ID)
		assert.ErrorIs(t, err, ErrProjectDoesNotExist)
	})
}
