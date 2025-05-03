package service

import (
	"context"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEntryService(t *testing.T) {

	t.Run("valid uow", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		require.NotPanics(t, func() {
			NewEntryService(uow)
		})
	})

	t.Run("invalid uow", func(t *testing.T) {
		require.Panics(t, func() {
			NewEntryService(nil)
		})
	})
}

func TestListEntries(t *testing.T) {
	project := domain.MakeProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	ctx := context.Background()

	t.Run("valid data", func(t *testing.T) {
		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}

		svc := NewEntryService(&uow)

		svc_entries, err := svc.ListEntries(ctx, project.ID)

		assert.Nil(t, err)
		assert.ElementsMatch(t, svc_entries, entries)
	})

	t.Run("empty data", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		svc := NewEntryService(uow)

		svc_entries, err := svc.ListEntries(ctx, project.ID)

		assert.Nil(t, err)
		assert.Equal(t, []*domain.Entry{}, svc_entries)
	})

	t.Run("repository error", func(t *testing.T) {
		entryRepo := &testutil.FailingEntryRepo{}
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		svc := NewEntryService(&uow)

		svc_entries, err := svc.ListEntries(ctx, project.ID)

		assert.Nil(t, svc_entries)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestCreateEntry(t *testing.T) {
	project := domain.MakeProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	ctx := context.Background()

	t.Run("valid data", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		svc := NewEntryService(&uow)
		input := CreateEntryInput{Body: "get a venue"}

		entry, err := svc.CreateEntry(ctx, project.ID, &input)

		assert.Nil(t, err)
		assert.Equal(t, input.Body, entry.Body)
	})

	t.Run("missing project", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		svc := NewEntryService(&uow)
		input := CreateEntryInput{Body: "get a venue"}

		non_existent_id := uuid.NewString()
		entry, err := svc.CreateEntry(ctx, non_existent_id, &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(non_existent_id))
	})

	t.Run("repository error", func(t *testing.T) {
		entryRepo := &testutil.FailingEntryRepo{}
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		svc := NewEntryService(&uow)
		input := CreateEntryInput{Body: "get a venue"}

		entry, err := svc.CreateEntry(ctx, project.ID, &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}
