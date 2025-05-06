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

func TestNewProjectAPIHanlder(t *testing.T) {
	t.Run("valid uow", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		require.NotPanics(t, func() {
			NewProjectAPIHandler(uow)
		})
	})

	t.Run("invalid uow", func(t *testing.T) {
		require.Panics(t, func() {
			NewProjectAPIHandler(nil)
		})
	})
}

func TestListProjects(t *testing.T) {
	t.Run("valid status ok", func(t *testing.T) {
		project := domain.NewProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		mux := http.NewServeMux()
		apiHandler := NewProjectAPIHandler(&uow)
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/projects/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty status ok", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewProjectAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/projects/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		project_repo := &testutil.FailingProjectRepo{}
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/projects/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetPoject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.NewProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.NewProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("new project", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewProjectAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"title": "Buy a horse"}`
		req := httptest.NewRequest(http.MethodPost, "/api/projects/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("invalid data", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewProjectAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"title": 1234}`
		req := httptest.NewRequest(http.MethodPost, "/api/projects/", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.NewProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"title": "Buy a horse"}`
		url := fmt.Sprintf("/api/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.NewProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"title": "Buy a horse"}`
		url := fmt.Sprintf("/api/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("invalid data", func(t *testing.T) {
		project := domain.NewProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"title": 1234}`
		url := fmt.Sprintf("/api/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		project_repo := in_memory.NewProjectRepository([]*domain.Project{project})
		uow := in_memory.InMemoryUnitOfWork{Projects: project_repo}
		apiHandler := NewProjectAPIHandler(&uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
