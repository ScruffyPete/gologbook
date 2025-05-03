package in_memory

import (
	"context"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryUnitOfWork struct {
	Users    domain.UserRepository
	Projects domain.ProjectReporitory
	Entries  domain.EntryRepository
}

func NewInMemoryUnitOfWork() *InMemoryUnitOfWork {
	return &InMemoryUnitOfWork{
		Users:    NewUserRepository(nil),
		Projects: NewProjectRepository(nil),
		Entries:  NewEntryRepository(nil),
	}
}

func (uow *InMemoryUnitOfWork) WithTx(
	_ context.Context,
	fn func(domain.RepoBundle) error,
) error {
	return fn(domain.RepoBundle{
		Projects: uow.Projects,
		Entries:  uow.Entries,
	})
}

func (uow *InMemoryUnitOfWork) Close() error {
	return nil
}
