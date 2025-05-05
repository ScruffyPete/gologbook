package in_memory

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListEntries(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)
		repo := NewEntryRepository(entries)

		repo_entries, err := repo.ListEntries(project.ID)

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_entries, entries)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entries := []*domain.Entry{}
		repo := NewEntryRepository(entries)

		repo_entries, err := repo.ListEntries(project.ID)

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_entries, entries)
	})
}

func TestCreateEntry(t *testing.T) {
	project := domain.NewProject("Hunt a boar")
	entry := domain.NewEntry(project.ID, "Get an axe")
	repo := NewEntryRepository([]*domain.Entry{entry})

	repo_entry, err := repo.CreateEntry(entry)

	assert.Nil(t, err)
	assert.Equal(t, entry, repo_entry)
}

func TestDeleteEntiries(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)
		repo := NewEntryRepository(entries)

		err := repo.DeleteEntries(project.ID)

		assert.Nil(t, err)
	})

	t.Run("missing project", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)
		repo := NewEntryRepository(entries)

		non_existent_id := uuid.NewString()
		err := repo.DeleteEntries(non_existent_id)

		assert.ErrorIs(t, err, domain.NewErrProjectDoesNotExist(non_existent_id))
	})
}
