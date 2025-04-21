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

func TestListProjects(t *testing.T) {
	t.Run("valid status ok", func(t *testing.T) {
		project := domain.MakeProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		req := httptest.NewRequest(http.MethodGet, "/projects", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty status ok", func(t *testing.T) {
		project_repo := in_memory.NewProjectRepository(nil)
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		req := httptest.NewRequest(http.MethodGet, "/projects", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		project_repo := &testutil.FailingProjectRepo{}
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		req := httptest.NewRequest(http.MethodGet, "/projects", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetPoject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		url := fmt.Sprintf("/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.MakeProject("Buid a shed")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		url := fmt.Sprintf("/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestCreateProject(t *testing.T) {
	t.Run("new project", func(t *testing.T) {
		project_repo := in_memory.NewProjectRepository(nil)
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		payload := `{"title": "Buy a horse"}`
		req := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("invalid data", func(t *testing.T) {
		project_repo := in_memory.NewProjectRepository(nil)
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		payload := `{"title": 1234}`
		req := httptest.NewRequest(http.MethodPost, "/projects", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		payload := `{"title": "Buy a horse"}`
		url := fmt.Sprintf("/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.MakeProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		payload := `{"title": "Buy a horse"}`
		url := fmt.Sprintf("/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("invalid data", func(t *testing.T) {
		project := domain.MakeProject("Cook a hog")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		payload := `{"title": 1234}`
		url := fmt.Sprintf("/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteProject(t *testing.T) {
	t.Run("valid project", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		url := fmt.Sprintf("/projects/%s", project.ID)
		req := httptest.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("invalid project", func(t *testing.T) {
		project := domain.MakeProject("Dig a hole")
		project_repo := in_memory.NewProjectRepository([]domain.Project{project})
		hand := NewHandler(project_repo)
		mux := hand.NewRouter()

		url := fmt.Sprintf("/projects/%s", uuid.NewString())
		req := httptest.NewRequest(http.MethodDelete, url, nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
