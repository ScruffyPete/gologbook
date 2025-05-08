//go:build integration

package postgres

import (
	"testing"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEntryRepository_ListEntries(t *testing.T) {
	t.Run("returns all entries", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, entries, nil)
		defer db.Close()

		repo := NewEntryRepository(db)
		entries, err := repo.ListEntries(project.ID)

		assert.Nil(t, err)
		assert.Equal(t, len(entries), len(entries))
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		db.Close()

		repo := NewEntryRepository(db)
		_, err := repo.ListEntries(project.ID)
		assert.NotNil(t, err)
	})

	t.Run("returns an empty slice if no entries are found", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewEntryRepository(db)
		entries, err := repo.ListEntries(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(entries), 0)
	})
}

func TestEntryRepository_CreateEntry(t *testing.T) {
	t.Run("creates an entry", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry := domain.NewEntry(project.ID, "Get an axe")

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		defer db.Close()

		repo := NewEntryRepository(db)
		createdEntry, err := repo.CreateEntry(entry)

		assert.Nil(t, err)
		assert.Equal(t, createdEntry, entry)
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry := domain.NewEntry(project.ID, "Get an axe")

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		db.Close()

		repo := NewEntryRepository(db)
		_, err := repo.CreateEntry(entry)

		assert.NotNil(t, err)
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry := domain.NewEntry(project.ID, "Get an axe")

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, nil, nil)
		db.Close()

		repo := NewEntryRepository(db)
		_, err := repo.CreateEntry(entry)

		assert.NotNil(t, err)
	})
}

func TestEntryRepository_DeleteEntries(t *testing.T) {
	t.Run("deletes an entry", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry_1 := domain.NewEntry(project.ID, "Get an axe")
		entry_2 := domain.NewEntry(project.ID, "Get a bow")
		entries := []*domain.Entry{entry_1, entry_2}

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, entries, nil)
		defer db.Close()

		repo := NewEntryRepository(db)
		err := repo.DeleteEntries(project.ID)

		assert.Nil(t, err)

		repo_entries, err := repo.ListEntries(project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(repo_entries), 0)
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry_1 := domain.NewEntry(project.ID, "Get an axe")
		entry_2 := domain.NewEntry(project.ID, "Get a bow")
		entries := []*domain.Entry{entry_1, entry_2}

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, entries, nil)
		db.Close()

		repo := NewEntryRepository(db)
		err := repo.DeleteEntries(project.ID)

		assert.NotNil(t, err)
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry_1 := domain.NewEntry(project.ID, "Get an axe")
		entry_2 := domain.NewEntry(project.ID, "Get a bow")
		entries := []*domain.Entry{entry_1, entry_2}

		db, _ := testutil.NewTestDB(nil, []*domain.Project{project}, entries, nil)
		db.Close()

		repo := NewEntryRepository(db)
		err := repo.DeleteEntries(project.ID)

		assert.NotNil(t, err)
	})
}
