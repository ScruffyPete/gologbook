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
