package in_memory

import (
	"sort"
	"testing"
	"time"

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
		baseTime := time.Now().UTC()
		older := &domain.Project{
			ID:        uuid.NewString(),
			CreatedAt: baseTime.Add(-time.Minute).Format(time.RFC3339Nano),
			Title:     "Old project",
		}
		newer := &domain.Project{
			ID:        uuid.NewString(),
			CreatedAt: baseTime.Format(time.RFC3339Nano),
			Title:     "New project",
		}
		projects := []*domain.Project{older, newer}
		sort.Slice(projects, func(i, j int) bool {
			return projects[i].CreatedAt > projects[j].CreatedAt
		})
		repo := NewProjectRepository(projects)

		repo_projects, err := repo.ListProjects()

		assert.Nil(t, err)
		assert.Equal(t, newer.Title, repo_projects[0].Title)
		assert.Equal(t, older.Title, repo_projects[1].Title)
	})
}

func TestGetProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.NewProject("Build a treehouse")
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
	project := domain.NewProject("Write a novel")
	repo := NewProjectRepository(nil)
	repo_project, err := repo.CreateProject(project)
	assert.Nil(t, err)
	assert.Equal(t, project, repo_project)
}

func TestUpdateProject(t *testing.T) {
	t.Run("existing project", func(t *testing.T) {
		project := domain.NewProject("Throw a ball")
		repo := NewProjectRepository([]*domain.Project{project})
		err := repo.UpdateProject(&domain.Project{ID: project.ID, Title: "Throw THE ball"})
		assert.ErrorIs(t, err, nil)
	})

	t.Run("missing project", func(t *testing.T) {
		project := domain.NewProject("Throw a ball")
		repo := NewProjectRepository(testutil.MakeDummyProjects())
		err := repo.UpdateProject(project)
		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(project.ID))
	})
}

func TestDeleteProject(t *testing.T) {
	project := domain.NewProject("Earn a million")

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
