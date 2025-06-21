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
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	project := domain.NewProject("Hunt a boar")
	entries := testutil.MakeDummyEntries(project)

	t.Run("returns all entries", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var gotEntries []*domain.Entry
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}
			for _, e := range entries {
				if _, err := repos.Entries.CreateEntry(ctx, e); err != nil {
					return err
				}
			}

			var err error
			gotEntries, err = repos.Entries.ListEntries(ctx, project.ID)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}

		assert.Equal(t, len(entries), len(gotEntries))
	})

	t.Run("returns an empty slice if no entries are found", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var gotEntries []*domain.Entry
		err := uow.WithTx(ctx, func(repos domain.RepoBundle) error {
			if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
				return err
			}

			gotEntries, err = repos.Entries.ListEntries(ctx, project.ID)
			return err
		})
		if err != nil {
			t.Fatalf("WithTx error: %v", err)
		}

		assert.Equal(t, 0, len(gotEntries))
	})
}

func TestEntryRepository_CreateEntry(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := domain.NewProject("Hunt a boar")
	entry := domain.NewEntry(project.ID, "Get an axe")

	var createdEntry *domain.Entry
	err = uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
			return err
		}
		createdEntry, err = repos.Entries.CreateEntry(ctx, entry)
		return err
	})
	if err != nil {
		t.Fatalf("WithTx error: %v", err)
	}
	assert.Equal(t, entry, createdEntry)
}

func TestEntryRepository_DeleteEntries(t *testing.T) {
	uow, err := NewTestUnitOfWork()
	if err != nil {
		t.Fatalf("failed to create unit of work: %v", err)
	}
	defer uow.Close()

	project := domain.NewProject("Hunt a boar")
	entry_1 := domain.NewEntry(project.ID, "Get an axe")
	entry_2 := domain.NewEntry(project.ID, "Get a bow")
	entries := []*domain.Entry{entry_1, entry_2}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var repoEntries []*domain.Entry
	err = uow.WithTx(ctx, func(repos domain.RepoBundle) error {
		if _, err := repos.Projects.CreateProject(ctx, project); err != nil {
			return err
		}
		for _, e := range entries {
			if _, err := repos.Entries.CreateEntry(ctx, e); err != nil {
				return err
			}
		}
		if err := repos.Entries.DeleteEntries(ctx, project.ID); err != nil {
			return err
		}
		repoEntries, err = repos.Entries.ListEntries(ctx, project.ID)
		return err
	})
	if err != nil {
		t.Fatalf("WithTx error: %v", err)
	}
	assert.Equal(t, len(repoEntries), 0)
}
