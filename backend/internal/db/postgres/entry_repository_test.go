//go:build integration

package postgres

import (
	"testing"
)

func TestEntryRepository_ListEntries(t *testing.T) {
	t.Run("returns all entries", func(t *testing.T) {
		return nil
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		return nil
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		return nil
	})

	t.Run("returns an empty slice if no entries are found", func(t *testing.T) {
		return nil
	})
}
