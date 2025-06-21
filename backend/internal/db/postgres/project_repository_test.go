//go:build integration

package postgres

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProjectRepository_ListProjects(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	older := domain.NewProject("Old project")
	newer := domain.NewProject("Newer project")
	newest := domain.NewProject("Newest project")
	projects := []*domain.Project{older, newer, newest}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].CreatedAt > projects[j].CreatedAt
	})

	t.Run("returns all projects", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var outputProjects []*domain.Project
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			for _, p := range projects {
				if _, err := repos.Projects.CreateProject(ctx, p); err != nil {
					return err
				}
			}
			outputProjects, err = repos.Projects.ListProjects(ctx)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.ElementsMatch(t, projects, outputProjects)
	})

	t.Run("returns an empty slice if no projects are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var outputProjects []*domain.Project
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			outputProjects, err = repos.Projects.ListProjects(ctx)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, len(outputProjects), 0)
	})
}

func TestProjectRepository_GetProject(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	project := domain.NewProject("Buy a farm")

	t.Run("returns a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var outputProject *domain.Project
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}
			outputProject, err = repos.Projects.GetProject(ctx, project.ID)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, outputProject, project)
	})

	t.Run("returns an error if the project does not exist", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			_, err = repos.Projects.GetProject(ctx, project.ID)
			return err
		})
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_CreateProject(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	project := domain.NewProject("Start a company")

	t.Run("creates a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var outputProject *domain.Project
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			outputProject, err = repos.Projects.CreateProject(ctx, project)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, project, outputProject)
	})

	t.Run("returns an error if the project already exists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}
			_, err := repos.Projects.CreateProject(ctx, project)
			return err
		})
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_UpdateProject(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}

	t.Run("updates a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")

		newTitle := "Start a company 2"
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}
			project.Title = newTitle
			return repos.Projects.UpdateProject(ctx, project)
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}
		assert.Equal(t, newTitle, project.Title)
	})

	t.Run("returns an error if the project does not exist", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")

		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			return repos.Projects.UpdateProject(ctx, project)
		})
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_DeleteProject(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := domain.NewProject("Start a company")

	err = uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
			return err
		}

		err := repos.Projects.DeleteProject(ctx, project.ID)
		return err
	})
	assert.Nil(t, err)
}
