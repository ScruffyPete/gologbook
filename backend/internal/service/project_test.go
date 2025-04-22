package service

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListProjects(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		projects := testutil.MakeDummyProjects()
		repo := in_memory.NewProjectRepository(projects)
		service := NewProjectService(repo)

		service_projects, err := service.ListProjects()

		assert.Nil(t, err)
		assert.ElementsMatch(t, projects, service_projects)
	})

	t.Run("empty data", func(t *testing.T) {
		repo := in_memory.NewProjectRepository(nil)
		service := NewProjectService(repo)

		service_projects, err := service.ListProjects()

		assert.Nil(t, err)
		assert.Equal(t, []*domain.Project{}, service_projects)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &testutil.FailingProjectRepo{}
		service := NewProjectService(repo)

		service_projects, err := service.ListProjects()

		assert.Nil(t, service_projects)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestGetProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		repo := in_memory.NewProjectRepository([]*domain.Project{project})
		service := NewProjectService(repo)

		service_project, err := service.GetProject(project.ID)

		assert.Nil(t, err)
		assert.Equal(t, project, service_project)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &testutil.FailingProjectRepo{}
		service := NewProjectService(repo)

		service_project, err := service.GetProject(uuid.NewString())

		assert.Nil(t, service_project)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("new project", func(t *testing.T) {
		repo := in_memory.NewProjectRepository(nil)
		service := NewProjectService(repo)

		title := "Buy a horse"
		input := CreateProjectInput{Title: title}

		project, err := service.CreateProject(&input)

		assert.Nil(t, err)
		assert.Equal(t, title, project.Title)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &testutil.FailingProjectRepo{}
		service := NewProjectService(repo)
		input := CreateProjectInput{Title: "Buy a horse"}

		project, err := service.CreateProject(&input)

		assert.Nil(t, project)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		repo := in_memory.NewProjectRepository([]*domain.Project{project})
		service := NewProjectService(repo)

		input := CreateProjectInput{Title: "Build a tree-fortress"}

		err := service.UpdateProject(project.ID, &input)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &testutil.FailingProjectRepo{}
		service := NewProjectService(repo)
		input := CreateProjectInput{Title: "Buy a horse"}

		err := service.UpdateProject(uuid.NewString(), &input)

		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		repo := in_memory.NewProjectRepository([]*domain.Project{project})
		service := NewProjectService(repo)

		err := service.DeleteProject(project.ID)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		repo := &testutil.FailingProjectRepo{}
		service := NewProjectService(repo)

		err := service.DeleteProject(uuid.NewString())

		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}
