package db

type ProjectReporitory interface {
	ListProjects() []Project
	GetProject(id string) *Project
	CreateProject(project Project) error
	UpdateProject(project Project) error
	DeleteProject(id string) error
}

type EntryRepository interface {
	ListEntries(projectID string) []Entry
	AddEntry(body string) error
}

func NewProjectRepository() ProjectReporitory {
	return newInMemoryProjectRepository(
		// FIXME
		map[string]Project{
			"project-1": {ID: "project-1", Title: "Project 1"},
			"project-2": {ID: "project-2", Title: "Project 2"},
			"project-3": {ID: "project-3", Title: "Project 3"},
		},
	)
}
