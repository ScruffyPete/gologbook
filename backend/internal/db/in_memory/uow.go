package in_memory

import (
	"context"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type InMemoryUnitOfWork struct {
	Users    domain.UserRepository
	Projects domain.ProjectReporitory
	Entries  domain.EntryRepository
	Insights domain.InsightRepository
}

func NewInMemoryUnitOfWork() *InMemoryUnitOfWork {
	return &InMemoryUnitOfWork{
		Users:    NewUserRepository(nil),
		Projects: NewProjectRepository(nil),
		Entries:  NewEntryRepository(nil),
		Insights: NewInsightRepository(nil),
	}
}

func (uow *InMemoryUnitOfWork) WithTx(
	_ context.Context,
	fn func(domain.RepoBundle) error,
) error {
	return fn(domain.RepoBundle{
		Users:    uow.Users,
		Projects: uow.Projects,
		Entries:  uow.Entries,
		Insights: uow.Insights,
	})
}

func (uow *InMemoryUnitOfWork) Close() error {
	return nil
}
