package testutil

import (
	"sort"
	"time"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/google/uuid"
)

func MakeDummyProjects() []*domain.Project {
	projectA := domain.NewProject("Build a treehouse")
	projectB := domain.NewProject("Paint the garage")
	projectC := domain.NewProject("Cook a feast")

	return []*domain.Project{
		projectA,
		projectB,
		projectC,
	}
}

func MakeDummyEntries(project *domain.Project) []*domain.Entry {
	entryA := domain.NewEntry(project.ID, "Get spear")
	entryB := domain.NewEntry(project.ID, "Build traps")
	entryC := domain.NewEntry(project.ID, "Sharpen knives")

	entries := []*domain.Entry{entryA, entryB, entryC}

	sorted := make([]*domain.Entry, len(entries))
	copy(sorted, entries)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CreatedAt > sorted[j].CreatedAt
	})

	return sorted
}

func NewDocument(projectID string, entryIDs []string, body string, createdAt *time.Time) *domain.Document {
	if createdAt == nil {
		t := time.Now().UTC()
		createdAt = &t
	}

	return &domain.Document{
		ID:        uuid.NewString(),
		CreatedAt: createdAt.Format("2006-01-02T15:04:05.999999Z"),
		ProjectID: projectID,
		EntryIDs:  entryIDs,
		Body:      body,
	}
}
