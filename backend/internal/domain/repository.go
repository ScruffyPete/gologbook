package domain

type ProjectReporitory interface {
	ListProjects() ([]*Project, error)
	GetProject(id string) (*Project, error)
	CreateProject(project *Project) (*Project, error)
	UpdateProject(project *Project) error
	DeleteProject(id string) error
}

type EntryRepository interface {
	ListEntries(projectID string) ([]*Entry, error)
	CreateEntry(entry *Entry) (*Entry, error)
	DeleteEntries(projectID string) error
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type InsightRepository interface {
	ListInsights(projectID string) ([]*Insight, error)
}

type RepoBundle struct {
	Users    UserRepository
	Projects ProjectReporitory
	Entries  EntryRepository
	Insights InsightRepository
}
