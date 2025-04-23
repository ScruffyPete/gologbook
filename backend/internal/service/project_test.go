package service

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProjectService(t *testing.T) {
	t.Run("valid uow", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		require.NotPanics(t, func() {
			NewProjectService(uow)
		})
	})

	t.Run("invalid uow", func(t *testing.T) {
		require.Panics(t, func() {
			NewProjectService(nil)
		})
	})
}

func TestListProjects(t *testing.T) {
	ctx := context.Background()

	t.Run("valid data", func(t *testing.T) {
		projects := testutil.MakeDummyProjects()
		uow := in_memory.InMemoryUnitOfWork{
			Projects: in_memory.NewProjectRepository(projects),
		}
		service := NewProjectService(&uow)

		service_projects, err := service.ListProjects(ctx)

		assert.Nil(t, err)
		assert.ElementsMatch(t, projects, service_projects)
	})

	t.Run("empty data", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		service := NewProjectService(uow)

		service_projects, err := service.ListProjects(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []*domain.Project{}, service_projects)
	})

	t.Run("repository error", func(t *testing.T) {
		uow := in_memory.InMemoryUnitOfWork{
			Projects: &testutil.FailingProjectRepo{},
		}
		service := NewProjectService(&uow)

		service_projects, err := service.ListProjects(ctx)

		assert.Nil(t, service_projects)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestGetProject(t *testing.T) {
	ctx := context.Background()

	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		uow := in_memory.InMemoryUnitOfWork{
			Projects: in_memory.NewProjectRepository([]*domain.Project{project}),
		}
		service := NewProjectService(&uow)

		service_project, err := service.GetProject(ctx, project.ID)

		assert.Nil(t, err)
		assert.Equal(t, project, service_project)
	})

	t.Run("repository error", func(t *testing.T) {
		uow := in_memory.InMemoryUnitOfWork{
			Projects: &testutil.FailingProjectRepo{},
		}
		service := NewProjectService(&uow)

		service_project, err := service.GetProject(ctx, uuid.NewString())

		assert.Nil(t, service_project)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestCreateProject(t *testing.T) {
	ctx := context.Background()

	t.Run("new project", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		service := NewProjectService(uow)

		title := "Buy a horse"
		input := CreateProjectInput{Title: title}

		project, err := service.CreateProject(ctx, &input)

		assert.Nil(t, err)
		assert.Equal(t, title, project.Title)
	})

	t.Run("repository error", func(t *testing.T) {
		uow := in_memory.InMemoryUnitOfWork{
			Projects: &testutil.FailingProjectRepo{},
		}
		service := NewProjectService(&uow)
		input := CreateProjectInput{Title: "Buy a horse"}

		project, err := service.CreateProject(ctx, &input)

		assert.Nil(t, project)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestUpdateProject(t *testing.T) {
	ctx := context.Background()
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		uow := in_memory.InMemoryUnitOfWork{
			Projects: in_memory.NewProjectRepository([]*domain.Project{project}),
		}
		service := NewProjectService(&uow)

		input := CreateProjectInput{Title: "Build a tree-fortress"}

		err := service.UpdateProject(ctx, project.ID, &input)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		uow := in_memory.InMemoryUnitOfWork{
			Projects: &testutil.FailingProjectRepo{},
		}
		service := NewProjectService(&uow)

		input := CreateProjectInput{Title: "Buy a horse"}

		err := service.UpdateProject(ctx, uuid.NewString(), &input)

		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestDeleteProject(t *testing.T) {
	ctx := context.Background()

	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Build a treehouse")
		uow := in_memory.InMemoryUnitOfWork{
			Projects: in_memory.NewProjectRepository([]*domain.Project{project}),
		}
		service := NewProjectService(&uow)

		err := service.DeleteProject(ctx, project.ID)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		uow := in_memory.InMemoryUnitOfWork{
			Projects: &testutil.FailingProjectRepo{},
		}
		service := NewProjectService(&uow)

		err := service.DeleteProject(ctx, uuid.NewString())

		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}
