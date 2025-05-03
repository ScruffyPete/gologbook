package in_memory

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
		repo := NewProjectRepository(projects)

		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_projects, projects)
	})

	t.Run("empty data", func(t *testing.T) {
		projects := []*domain.Project{}
		repo := NewProjectRepository(projects)
		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_projects, projects)
	})

	t.Run("ordered by CreatedAt", func(t *testing.T) {
		older := domain.MakeProject("Old project")
		newer := domain.MakeProject("New project")
		repo := NewProjectRepository([]*domain.Project{older, newer})

		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.Equal(t, newer.Title, repo_projects[0].Title)
		assert.Equal(t, older.Title, repo_projects[1].Title)
	})
}

func TestGetProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		repo := NewProjectRepository([]*domain.Project{project})
		repo_project, err := repo.GetProject(project.ID)
		assert.Equal(t, repo_project, project)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("invalid project", func(t *testing.T) {
		repo := NewProjectRepository(testutil.MakeDummyProjects())
		non_existent_id := uuid.NewString()
		repo_project, err := repo.GetProject(non_existent_id)
		assert.Nil(t, repo_project)
		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(non_existent_id))
	})
}

func TestCreateProject(t *testing.T) {
	project := domain.MakeProject("Write a novel")
	repo := NewProjectRepository(nil)
	repo_project, err := repo.CreateProject(project)
	assert.Nil(t, err)
	assert.Equal(t, project, repo_project)
}

func TestUpdateProject(t *testing.T) {
	t.Run("existing project", func(t *testing.T) {
		project := domain.MakeProject("Throw a ball")
		repo := NewProjectRepository([]*domain.Project{project})
		err := repo.UpdateProject(&domain.Project{ID: project.ID, Title: "Throw THE ball"})
		assert.ErrorIs(t, err, nil)
	})

	t.Run("missing project", func(t *testing.T) {
		project := domain.MakeProject("Throw a ball")
		repo := NewProjectRepository(testutil.MakeDummyProjects())
		err := repo.UpdateProject(project)
		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(project.ID))
	})
}

func TestDeleteProject(t *testing.T) {
	project := domain.MakeProject("Earn a million")

	t.Run("existing project", func(t *testing.T) {
		repo := NewProjectRepository([]*domain.Project{project})
		err := repo.DeleteProject(project.ID)
		assert.ErrorIs(t, err, nil)
	})

	t.Run("invalid project", func(t *testing.T) {
		repo := NewProjectRepository(testutil.MakeDummyProjects())
		err := repo.DeleteProject(project.ID)
		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(project.ID))
	})
}
