package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDocumentAPIHandler_ListDocuments(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entry := domain.NewEntry(project.ID, "Dig a hole in the hole")
		document := testutil.NewDocument(project.ID, []string{entry.ID}, "Test Document", nil)
		documentRepo := in_memory.NewDocumentRepository([]*domain.Document{document})
		uow := in_memory.InMemoryUnitOfWork{
			Documents: documentRepo,
		}
		apiHandler := NewDocumentAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/documents/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		documentRepo := in_memory.NewDocumentRepository([]*domain.Document{})
		uow := in_memory.InMemoryUnitOfWork{
			Projects:  projectRepo,
			Documents: documentRepo,
		}
		apiHandler := NewDocumentAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/documents/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("missing project id", func(t *testing.T) {
		documentRepo := in_memory.NewDocumentRepository([]*domain.Document{})
		uow := in_memory.InMemoryUnitOfWork{
			Documents: documentRepo,
		}
		apiHandler := NewDocumentAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/documents/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
