package postgres

import (
	"context"
	"database/sql"
	"os"

	"github.com/ScruffyPete/gologbook/internal/domain"
	_ "github.com/lib/pq"
)

type PostgresUnitOfWork struct {
	db *sql.DB
}

func NewPostgresUnitOfWork() (*PostgresUnitOfWork, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &PostgresUnitOfWork{db: db}, nil
}

func (uow *PostgresUnitOfWork) WithTx(ctx context.Context, fn func(repos domain.RepoBundle) error) error {
	tx, err := uow.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repos := domain.RepoBundle{
		Users:     NewUserRepository(uow.db),
		Projects:  NewProjectRepository(uow.db),
		Entries:   NewEntryRepository(uow.db),
		Documents: NewDocumentRepository(uow.db),
	}

	if err := fn(repos); err != nil {
		return err
	}

	return tx.Commit()
}

func (uow *PostgresUnitOfWork) Close() error {
	return uow.db.Close()
}
