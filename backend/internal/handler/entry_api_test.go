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
	"github.com/stretchr/testify/require"
)

func TestNewEntryAPIHanlder(t *testing.T) {
	t.Run("valid uow", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		require.NotPanics(t, func() {
			NewEntryAPIHandler(uow)
		})
	})

	t.Run("invalid uow", func(t *testing.T) {
		require.Panics(t, func() {
			NewEntryAPIHandler(nil)
		})
	})
}

func TestListEntries(t *testing.T) {

	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)
		uow := in_memory.InMemoryUnitOfWork{
			Entries: entryRepo,
		}

		apiHandler := NewAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("reposirotry error", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entryRepo := &testutil.FailingEntryRepo{}
		uow := in_memory.InMemoryUnitOfWork{
			Entries: entryRepo,
		}
		apiHandler := NewAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/{%s}/entries", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCreateEntry(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		entryRepo := in_memory.NewEntryRepository(nil)
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  entryRepo,
		}
		apiHandler := NewAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"body": "Get a shovel"}`
		url := fmt.Sprintf("/api/projects/%s/entries", project.ID)
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"body": "Get a shovel"}`
		url := fmt.Sprintf("/api/projects/%s/entries", uuid.NewString())
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("invalid input", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"body": 123}`
		url := fmt.Sprintf("/api/projects/%s/entries", project.ID)
		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
