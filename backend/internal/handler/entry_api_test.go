package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListEntries(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		url := fmt.Sprintf("/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entryRepo := in_memory.NewEntryRepository(nil)
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		url := fmt.Sprintf("/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("reposirotry error", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entryRepo := &testutil.FailingEntryRepo{}
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		url := fmt.Sprintf("/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCreateEntry(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entryRepo := in_memory.NewEntryRepository(nil)
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		payload := `{"body": "Get a shovel"}`
		url := fmt.Sprintf("/projects/%s/entries", project.ID)
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entryRepo := in_memory.NewEntryRepository(nil)
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		payload := `{"body": "Get a shovel"}`
		url := fmt.Sprintf("/projects/%s/entries", uuid.NewString())
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("invalid input", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		projectHandler := NewProjectHandler(projectRepo)

		entryRepo := in_memory.NewEntryRepository(nil)
		entryHandler := NewEntryHandler(entryRepo, projectRepo)

		mux := http.NewServeMux()
		projectHandler.Register(mux)
		entryHandler.Register(mux)

		payload := `{"body": 123}`
		url := fmt.Sprintf("/projects/%s/entries", project.ID)
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
