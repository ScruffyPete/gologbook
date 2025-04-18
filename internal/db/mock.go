package db

import "fmt"

type mockProjectRepository struct {
	projects map[string]Project
}

func (repo *mockProjectRepository) ListProjects() []Project {
	projects := make([]Project, 0, len(repo.projects))

	for _, project := range repo.projects {
		projects = append(projects, project)
	}

	return projects
}

func (repo *mockProjectRepository) GetProjectByID(id string) *Project {
	projectData, ok := repo.projects[id]
	if !ok {
		return nil
	}

	return &projectData
}

func (repo *mockProjectRepository) CreateProject(title string) error {
	projectCount := len(repo.projects)
	id := fmt.Sprintf("project-%d", projectCount+1)
	project := Project{
		ID:    id,
		Title: title,
	}
	repo.projects[id] = project
	return nil
}

func (repo *mockProjectRepository) LogEntry(id string) error {
	return nil // TODO
}

func (repo *mockProjectRepository) DeleteProject(id string) error {
	delete(repo.projects, id)
	return nil
}

func (repo *mockProjectRepository) SetupRepository() error {
	repo.projects = map[string]Project{
		"project-1": {
			ID:    "project-1",
			Title: "Project 1",
		},
		"project-2": {
			ID:    "project-2",
			Title: "Project 2",
		},
		"project-3": {
			ID:    "project-3",
			Title: "Project 3",
		},
	}
	return nil
}
