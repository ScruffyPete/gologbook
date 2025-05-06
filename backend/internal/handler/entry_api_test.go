package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/queue"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListEntries(t *testing.T) {

	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entries := testutil.MakeDummyEntries(project)
		entryRepo := in_memory.NewEntryRepository(entries)
		uow := in_memory.InMemoryUnitOfWork{
			Entries: entryRepo,
		}
		apiHandler := NewEntryAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/entries/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Entries:  in_memory.NewEntryRepository(nil),
		}
		apiHandler := NewEntryAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/entries/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("missing project id", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
		}
		apiHandler := NewEntryAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/entries/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("reposirotry error", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entryRepo := &testutil.FailingEntryRepo{}
		uow := in_memory.InMemoryUnitOfWork{
			Entries: entryRepo,
		}
		apiHandler := NewEntryAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/entries/?project_id=%s", project.ID)
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
		queue := queue.NewInMemoryQueue()
		apiHandler := NewEntryAPIHandler(&uow, queue)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := fmt.Sprintf(`{"body": "Get a shovel", "project_id": "%s"}`, project.ID)
		req := httptest.NewRequest(http.MethodPost, "/api/entries/", strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		queue := queue.NewInMemoryQueue()
		apiHandler := NewEntryAPIHandler(uow, queue)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := fmt.Sprintf(`{"body": "Get a shovel", "project_id": "%s"}`, uuid.NewString())
		req := httptest.NewRequest(http.MethodPost, "/api/entries/", strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("invalid input", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		uow := &in_memory.InMemoryUnitOfWork{
			Projects: in_memory.NewProjectRepository([]*domain.Project{project}),
			Entries:  in_memory.NewEntryRepository(nil),
		}
		queue := queue.NewInMemoryQueue()
		apiHandler := NewEntryAPIHandler(uow, queue)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := fmt.Sprintf(`{"body": 123, "project_id": "%s"}`, project.ID)
		req := httptest.NewRequest(http.MethodPost, "/api/entries/", strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("queue error", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		uow := in_memory.NewInMemoryUnitOfWork()
		queue := &testutil.FailingQueue{}
		apiHandler := NewEntryAPIHandler(uow, queue)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := fmt.Sprintf(`{"body": "Get a shovel", "project_id": "%s"}`, project.ID)
		req := httptest.NewRequest(http.MethodPost, "/api/entries/", strings.NewReader(payload))
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
