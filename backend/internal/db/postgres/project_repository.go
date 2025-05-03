package postgres

import (
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *projectRepository {
	return &projectRepository{db: db}
}

func (repo *projectRepository) ListProjects() ([]*domain.Project, error) {
	rows, err := repo.db.Query("SELECT id, title, created_at FROM projects ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]*domain.Project, 0)
	for rows.Next() {
		var p domain.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	return projects, nil
}

func (repo *projectRepository) GetProject(id string) (*domain.Project, error) {
	row, err := repo.db.Query("SELECT id, title, created_at FROM projects WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, domain.NewErrProjectDoesNotExist(id)
	}

	var p domain.Project
	if err := row.Scan(&p.ID, &p.Title, &p.CreatedAt); err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *projectRepository) CreateProject(project *domain.Project) (*domain.Project, error) {
	return nil, nil
}

func (repo *projectRepository) UpdateProject(project *domain.Project) error {
	return nil
}

func (repo *projectRepository) DeleteProject(id string) error {
	return nil
}
