package db

import "github.com/ScruffyPete/gologbook/api"

type ProjectReporitory interface {
	ListProjects() []Project
	GetProject(id string) *Project
	CreateProject(title string) error
	UpdateProject(id string, updates *api.ProjectRequestBody) error
	DeleteProject(id string) error

	init() error
}

type EntryRepository interface {
	ListEntries(projectID string) []Entry
	AddEntry(body string) error

	init() error
}

func NewProjectRepository() ProjectReporitory {
	projectRepo := &mockProjectRepository{}
	if err := projectRepo.init(); err != nil {
		panic("Failed to init project repository") // TODO internal external error handling
	}
	return projectRepo
}
