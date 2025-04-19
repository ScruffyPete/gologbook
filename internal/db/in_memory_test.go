package db

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func makeProject(title string) Project {
	return Project{
		ID:    uuid.NewString(),
		Title: title,
	}
}

func makeProjectMap() map[string]Project {
	projectA := makeProject("Build a treehouse")
	projectB := makeProject("Paint the garage")
	projectC := makeProject("Cook a feast")

	return map[string]Project{
		projectA.ID: projectA,
		projectB.ID: projectB,
		projectC.ID: projectC,
	}
}

func TestListProjects(t *testing.T) {
	testCases := []struct {
		name     string
		data     map[string]Project
		expected []Project
	}{
		{
			name: "valid data",
			data: makeProjectMap(),
		},
		{
			name: "empty data",
			data: map[string]Project{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			repo := newInMemoryProjectRepository(tc.data)
			projects := repo.ListProjects()

			expected := make([]Project, 0, len(tc.data))
			for _, p := range tc.data {
				expected = append(expected, p)
			}

			assert.ElementsMatch(t, expected, projects)
		})
	}
}

func TestGetProject(t *testing.T) {
	project := makeProject("Build a treehouse")

	testCases := []struct {
		name     string
		data     map[string]Project
		id       string
		expected *Project
	}{
		{
			name:     "valid project",
			data:     map[string]Project{project.ID: project},
			id:       project.ID,
			expected: &project,
		},
		{
			name:     "invalid project",
			data:     makeProjectMap(),
			id:       uuid.NewString(),
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			repo := newInMemoryProjectRepository(tc.data)
			project := repo.GetProject(tc.id)
			assert.Equal(t, tc.expected, project)
		})
	}
}

func TestCreateProject(t *testing.T) {
	project := makeProject("Write a novel")

	testCases := []struct {
		name    string
		data    map[string]Project
		project Project
		err     error
	}{
		{
			name:    "new project",
			data:    makeProjectMap(),
			project: project,
			err:     nil,
		},
		{
			name:    "existing project",
			data:    map[string]Project{project.ID: project},
			project: project,
			err:     ErrProjectExists,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			repo := newInMemoryProjectRepository(tc.data)
			err := repo.CreateProject(tc.project)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestUpdateProject(t *testing.T) {
	project := makeProject("Throw a ball")

	testCases := []struct {
		name    string
		data    map[string]Project
		project Project
		err     error
	}{
		{
			name:    "existing project",
			data:    map[string]Project{project.ID: project},
			project: Project{ID: project.ID, Title: "Throw THE ball"},
			err:     nil,
		},
		{
			name:    "missing project",
			data:    makeProjectMap(),
			project: project,
			err:     ErrProjectDoesNotExist,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			repo := newInMemoryProjectRepository(tc.data)
			err := repo.UpdateProject(tc.project)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestDeleteProject(t *testing.T) {
	project := makeProject("Earn a million")

	testCases := []struct {
		name string
		data map[string]Project
		id   string
		err  error
	}{
		{
			name: "existing project",
			data: map[string]Project{project.ID: project},
			id:   project.ID,
			err:  nil,
		},
		{
			name: "missing project",
			data: makeProjectMap(),
			id:   project.ID,
			err:  ErrProjectDoesNotExist,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			repo := newInMemoryProjectRepository(tc.data)
			err := repo.DeleteProject(tc.id)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}
