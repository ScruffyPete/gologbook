package postgres

import (
	"context"
	"database/sql"
	"os"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type TestUnitOfWork struct {
	db *sql.DB
}

func NewTestUnitOfWork() (*TestUnitOfWork, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &TestUnitOfWork{db: db}, nil
}

func (uow *TestUnitOfWork) WithTx(ctx context.Context, fn func(repos domain.RepoBundle) error) error {
	tx, err := uow.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	repos := domain.RepoBundle{
		Users:     NewUserRepository(tx),
		Projects:  NewProjectRepository(tx),
		Entries:   NewEntryRepository(tx),
		Documents: NewDocumentRepository(tx),
	}
	return fn(repos)
}

func (uow *TestUnitOfWork) Close() error {
	return uow.db.Close()
}
