package service

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListEntries(t *testing.T) {
	project := domain.MakeProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})

	t.Run("valid data", func(t *testing.T) {
		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)

		svc := NewEntryService(entryRepo, projectRepo)

		svc_entries, err := svc.ListEntries(project.ID)

		assert.Nil(t, err)
		assert.ElementsMatch(t, svc_entries, entries)
	})

	t.Run("empty data", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		svc := NewEntryService(entryRepo, projectRepo)

		svc_entries, err := svc.ListEntries(project.ID)

		assert.Nil(t, err)
		assert.Equal(t, []*domain.Entry{}, svc_entries)
	})

	t.Run("repository error", func(t *testing.T) {
		entryRepo := &testutil.FailingEntryRepo{}
		svc := NewEntryService(entryRepo, projectRepo)

		svc_entries, err := svc.ListEntries(project.ID)

		assert.Nil(t, svc_entries)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}

func TestCreateEntry(t *testing.T) {
	project := domain.MakeProject("Plan a wedding")
	projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})

	t.Run("valid data", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		svc := NewEntryService(entryRepo, projectRepo)
		input := CreateEntryInput{Body: "get a venue"}

		entry, err := svc.CreateEntry(project.ID, &input)

		assert.Nil(t, err)
		assert.Equal(t, input.Body, entry.Body)
	})

	t.Run("missing project", func(t *testing.T) {
		entryRepo := in_memory.NewEntryRepository(nil)
		svc := NewEntryService(entryRepo, projectRepo)
		input := CreateEntryInput{Body: "get a venue"}

		entry, err := svc.CreateEntry(uuid.NewString(), &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, in_memory.ErrProjectDoesNotExist)
	})

	t.Run("repository error", func(t *testing.T) {
		entryRepo := &testutil.FailingEntryRepo{}
		svc := NewEntryService(entryRepo, projectRepo)
		input := CreateEntryInput{Body: "get a venue"}

		entry, err := svc.CreateEntry(project.ID, &input)

		assert.Nil(t, entry)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, testutil.ErrRepoFailed)
	})
}
