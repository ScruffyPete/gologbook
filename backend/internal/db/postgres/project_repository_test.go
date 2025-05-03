//go:build integration

package postgres

import (
	"testing"
)

func TestProjectRepository_ListProjects(t *testing.T) {
	t.Run("returns all projects", func(t *testing.T) {
		return nil
	})

	t.Run("returns an error if the database connection fails", func(t *testing.T) {
		return nil
	})

	t.Run("returns an error if the query fails", func(t *testing.T) {
		return nil
	})

	t.Run("returns an empty slice if no projects are found", func(t *testing.T) {
		return nil
	})
}
