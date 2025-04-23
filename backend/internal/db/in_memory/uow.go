package in_memory

import (
	"context"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryUnitOfWork struct {
	Projects domain.ProjectReporitory
	Entries  domain.EntryRepository
}

func NewInMemoryUnitOfWork() *InMemoryUnitOfWork {
	return &InMemoryUnitOfWork{
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
