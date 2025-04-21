package testutil

import (
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

	return []*domain.Entry{entryA, entryB, entryC}
}
