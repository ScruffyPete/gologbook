package in_memory

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestListEntries(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.MakeProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)
		repo := NewEntryRepository(entries)

		repo_entries, err := repo.ListEntries()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_entries, entries)
	})

	t.Run("empty data", func(t *testing.T) {
		entries := []*domain.Entry{}
		repo := NewEntryRepository(entries)

		repo_entries, err := repo.ListEntries()

		assert.Nil(t, err)
		assert.ElementsMatch(t, repo_entries, entries)
	})
}

func TestCreateEntry(t *testing.T) {
	project := domain.MakeProject("Hunt a boar")
	entry := domain.MakeEntry(project.ID, "Get an axe")
	repo := NewEntryRepository([]*domain.Entry{entry})

	err := repo.CreateEntry(entry)

	assert.Nil(t, err)
}
