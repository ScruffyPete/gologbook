//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEntryRepository_ListEntries(t *testing.T) {
	t.Run("returns all entries", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Hunt a boar")
		entries := testutil.MakeDummyEntries(project)

		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, entries, nil)

		repo := NewEntryRepository(db)
		entries, err := repo.ListEntries(ctx, project.ID)

		assert.Nil(t, err)
		assert.Equal(t, len(entries), len(entries))
	})

	t.Run("returns an empty slice if no entries are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Hunt a boar")
		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, nil, nil)

		repo := NewEntryRepository(db)
		entries, err := repo.ListEntries(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(entries), 0)
	})
}

func TestEntryRepository_CreateEntry(t *testing.T) {
	t.Run("creates an entry", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		project := domain.NewProject("Hunt a boar")
		entry := domain.NewEntry(project.ID, "Get an axe")

		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, nil, nil)

		repo := NewEntryRepository(db)
		createdEntry, err := repo.CreateEntry(ctx, entry)

		assert.Nil(t, err)
		assert.Equal(t, createdEntry, entry)
	})
}

func TestEntryRepository_DeleteEntries(t *testing.T) {
	t.Run("deletes an entry", func(t *testing.T) {
		project := domain.NewProject("Hunt a boar")
		entry_1 := domain.NewEntry(project.ID, "Get an axe")
		entry_2 := domain.NewEntry(project.ID, "Get a bow")
		entries := []*domain.Entry{entry_1, entry_2}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db := testutil.NewTestDB(t, ctx, nil, []*domain.Project{project}, entries, nil)

		repo := NewEntryRepository(db)
		err := repo.DeleteEntries(ctx, project.ID)

		assert.Nil(t, err)

		repo_entries, err := repo.ListEntries(ctx, project.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(repo_entries), 0)
	})
}
