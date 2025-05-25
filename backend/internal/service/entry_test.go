package service

import (
	"context"
	"os"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/queue"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEntryService(t *testing.T) {

	t.Run("valid uow and queue", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		queue := queue.NewInMemoryQueue()
		require.NotPanics(t, func() {
			NewEntryService(uow, queue)
		})
	})

	t.Run("invalid uow and queue", func(t *testing.T) {
		require.Panics(t, func() {
			NewEntryService(nil, nil)
		})
	})

	t.Run("invalid uow", func(t *testing.T) {
		require.Panics(t, func() {
			NewEntryService(nil, nil)
		})
	})

	t.Run("invalid queue", func(t *testing.T) {
		require.Panics(t, func() {
			NewEntryService(nil, nil)
		})
	})
}

func TestListEntries(t *testing.T) {
	project := domain.NewProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	ctx := context.Background()

	t.Run("valid data", func(t *testing.T) {
		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(&uow, queue)

		svc_entries, err := svc.ListEntries(ctx, project.ID)

		assert.Nil(t, err)
		assert.ElementsMatch(t, svc_entries, entries)
	})

	t.Run("empty data", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(uow, queue)

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
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(&uow, queue)

		svc_entries, err := svc.ListEntries(ctx, project.ID)

		assert.Empty(t, svc_entries)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestCreateEntry(t *testing.T) {
	project := domain.NewProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
	ctx := context.Background()

	t.Run("valid data", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := &in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(uow, queue)
		input := CreateEntryInput{Body: "get a venue", ProjectID: project.ID}

		entry, err := svc.CreateEntry(ctx, &input)

		assert.Nil(t, err)
		assert.Equal(t, input.Body, entry.Body)

		key := os.Getenv("REDIS_PENDING_PROJECTS_KEY")
		timestamp, err := queue.Pop(key, entry.ProjectID)
		assert.Nil(t, err)
		assert.NotZero(t, timestamp)
	})

	t.Run("missing project", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(&uow, queue)
		non_existent_id := uuid.NewString()
		input := CreateEntryInput{Body: "get a venue", ProjectID: non_existent_id}

		entry, err := svc.CreateEntry(ctx, &input)

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
		queue := queue.NewInMemoryQueue()
		svc := NewEntryService(&uow, queue)
		input := CreateEntryInput{Body: "get a venue", ProjectID: project.ID}

		entry, err := svc.CreateEntry(ctx, &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})

	t.Run("queue error", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		queue := &testutil.FailingQueue{}
		svc := NewEntryService(&uow, queue)
		input := CreateEntryInput{Body: "get a venue", ProjectID: project.ID}

		entry, err := svc.CreateEntry(ctx, &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrQueueFailed)
	})
}
