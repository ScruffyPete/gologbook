package domain

import "context"

type ProjectReporitory interface {
	ListProjects(ctx context.Context) ([]*Project, error)
	GetProject(ctx context.Context, id string) (*Project, error)
	CreateProject(ctx context.Context, project *Project) (*Project, error)
	UpdateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, id string) error
}

type EntryRepository interface {
	ListEntries(ctx context.Context, projectID string) ([]*Entry, error)
	CreateEntry(ctx context.Context, entry *Entry) (*Entry, error)
	DeleteEntries(ctx context.Context, projectID string) error
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type DocumentRepository interface {
	ListDocuments(projectID string) ([]*Document, error)
}

type RepoBundle struct {
	Users     UserRepository
	Projects  ProjectReporitory
	Entries   EntryRepository
	Documents DocumentRepository
}
