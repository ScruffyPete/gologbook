package postgres

import (
	"context"
	"database/sql"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type projectRepository struct {
	tx *sql.Tx
}

func NewProjectRepository(tx *sql.Tx) *projectRepository {
	return &projectRepository{tx: tx}
}

func (repo *projectRepository) ListProjects(ctx context.Context) ([]*domain.Project, error) {
	rows, err := repo.tx.QueryContext(
		ctx,
		"SELECT id, title, created_at FROM projects ORDER BY created_at DESC",
	)
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

func (repo *projectRepository) GetProject(ctx context.Context, id string) (*domain.Project, error) {
	row, err := repo.tx.QueryContext(
		ctx,
		"SELECT id, title, created_at FROM projects WHERE id = $1",
		id,
	)
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

func (repo *projectRepository) CreateProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	_, err := repo.tx.ExecContext(
		ctx,
		"INSERT INTO projects (id, title, created_at) VALUES ($1, $2, $3)",
		project.ID,
		project.Title,
		project.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (repo *projectRepository) UpdateProject(ctx context.Context, project *domain.Project) error {
	result, err := repo.tx.ExecContext(
		ctx,
		"UPDATE projects SET title = $1 WHERE id = $2",
		project.Title,
		project.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.NewErrProjectDoesNotExist(project.ID)
	}

	return nil
}

func (repo *projectRepository) DeleteProject(ctx context.Context, id string) error {
	result, err := repo.tx.ExecContext(ctx, "DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.NewErrProjectDoesNotExist(id)
	}

	return nil
}
