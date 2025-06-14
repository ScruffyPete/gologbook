//go:build integration

package postgres

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestProjectRepository_ListProjects(t *testing.T) {
	t.Run("returns all projects", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		older := domain.NewProject("Old project")
		newer := domain.NewProject("Newer project")
		newest := domain.NewProject("Newest project")
		projects := []*domain.Project{older, newer, newest}
		sort.Slice(projects, func(i, j int) bool {
			return projects[i].CreatedAt > projects[j].CreatedAt
		})

		db, _ := testutil.NewTestDB(ctx, nil, projects, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		repo_projects, err := repo.ListProjects(ctx)

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_projects, projects)
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		db.Close() // Close immediately to force an error
		repo := NewProjectRepository(db)

		_, err := repo.ListProjects(ctx)
		assert.NotNil(t, err)
	})

	t.Run("returns an empty slice if no projects are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		defer db.Close()
		repo := NewProjectRepository(db)

		repo_projects, err := repo.ListProjects(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(repo_projects), 0)
	})
}

func TestProjectRepository_GetProject(t *testing.T) {
	t.Run("returns a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Buy a farm")
		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		repo_project, err := repo.GetProject(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, repo_project, project)
	})

	t.Run("returns an error if the project does not exist", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		_, err := repo.GetProject(ctx, "non-existent-id")
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_CreateProject(t *testing.T) {
	t.Run("creates a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")
		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		repo_project, err := repo.CreateProject(ctx, project)
		assert.Nil(t, err)
		assert.Equal(t, repo_project, project)
	})

	t.Run("returns an error if the project already exists", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")
		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, nil)
		defer db.Close()
		repo := NewProjectRepository(db)
		_, err := repo.CreateProject(ctx, project)
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_UpdateProject(t *testing.T) {
	t.Run("updates a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")
		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		project.Title = "Start a company 2"

		err := repo.UpdateProject(ctx, project)
		assert.Nil(t, err)

		repo_project, err := repo.GetProject(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, repo_project.Title, project.Title)
	})

	t.Run("returns an error if the project does not exist", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")
		db, _ := testutil.NewTestDB(ctx, nil, nil, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		project.Title = "Start a company 2"

		err := repo.UpdateProject(ctx, project)
		assert.NotNil(t, err)
	})
}

func TestProjectRepository_DeleteProject(t *testing.T) {
	t.Run("deletes a project", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		project := domain.NewProject("Start a company")
		db, _ := testutil.NewTestDB(ctx, nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewProjectRepository(db)
		err := repo.DeleteProject(ctx, project.ID)
		assert.Nil(t, err)

		repo_project, err := repo.GetProject(ctx, project.ID)
		assert.NotNil(t, err)
		assert.Nil(t, repo_project)
	})
}
