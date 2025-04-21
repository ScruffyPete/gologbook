package domain

type ProjectReporitory interface {
	ListProjects() []Project
	GetProject(id string) (*Project, error)
	CreateProject(project Project) error
	UpdateProject(project Project) error
	DeleteProject(id string) error
}

type EntryRepository interface {
	ListEntries(projectID string) []Entry
	AddEntry(body string) error
}
