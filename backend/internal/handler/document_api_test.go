package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/queue"
	"github.com/ScruffyPete/gologbook/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDocumentAPIHandler_GetLatestDocument(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		entry := domain.NewEntry(project.ID, "Dig a hole in the hole")
		document := testutil.NewDocument(project.ID, []string{entry.ID}, "Test Document", nil)
		documentRepo := in_memory.NewDocumentRepository([]*domain.Document{document})
		uow := in_memory.InMemoryUnitOfWork{
			Projects:  projectRepo,
			Documents: documentRepo,
		}
		apiHandler := NewDocumentAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/documents/%s/output/", project.ID)
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
		apiHandler := NewDocumentAPIHandler(&uow, nil)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/documents/%s/output/", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDocumentAPIHandler_StreamDocument(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		projectID := uuid.NewString()
		q := queue.NewInMemoryQueue()

		ch := make(chan string, 3)
		ch <- "first message"
		ch <- "second message"
		ch <- "[[STOP]]"
		close(ch)
		q.SetDocumentStream(projectID, ch)

		apiHandler := NewDocumentAPIHandler(nil, q)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/documents/%s/stream/", projectID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/event-stream", w.Header().Get("Content-Type"))
		assert.Equal(t, "no-cache", w.Header().Get("Cache-Control"))
		assert.Equal(t, "keep-alive", w.Header().Get("Connection"))

		data, err := io.ReadAll(w.Result().Body)
		assert.Nil(t, err)

		expectedData := "data: first message\n\ndata: second message\n\n"
		assert.Equal(t, expectedData, string(data))
	})
}
