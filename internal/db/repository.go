package db

type ProjectReporitory interface {
	ListProjects() []Project
	GetProjectByID(id string) *Project
	CreateProject(title string) error
	LogEntry(id string) error
	DeleteProject(id string) error

	SetupRepository() error
}

func NewRepository() (ProjectReporitory, error) {
	var repo ProjectReporitory = &mockProjectRepository{}

	if err := repo.SetupRepository(); err != nil {
		return nil, err
	}

	return repo, nil
}
