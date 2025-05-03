package testutil

import (
	"sort"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

func MakeDummyProjects() []*domain.Project {
	projectA := domain.MakeProject("Build a treehouse")
	projectB := domain.MakeProject("Paint the garage")
	projectC := domain.MakeProject("Cook a feast")

	return []*domain.Project{
		projectA,
		projectB,
		projectC,
	}
}

func MakeDummyEntries(project *domain.Project) []*domain.Entry {
	entryA := domain.MakeEntry(project.ID, "Get spear")
	entryB := domain.MakeEntry(project.ID, "Build traps")
	entryC := domain.MakeEntry(project.ID, "Sharpen knives")

	entries := []*domain.Entry{entryA, entryB, entryC}

	sorted := make([]*domain.Entry, len(entries))
	copy(sorted, entries)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CreatedAt > sorted[j].CreatedAt
	})

	return sorted
}
