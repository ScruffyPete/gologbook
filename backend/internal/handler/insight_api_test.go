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

func TestInsightAPIHandler_ListInsights(t *testing.T) {
	t.Run("valid data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		entry := domain.NewEntry(project.ID, "Dig a hole in the hole")
		insight := testutil.NewInsight(project.ID, []string{entry.ID}, "Test Insight", nil)
		insightRepo := in_memory.NewInsightRepository([]*domain.Insight{insight})
		uow := in_memory.InMemoryUnitOfWork{
			Insights: insightRepo,
		}
		apiHandler := NewInsightAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/insights/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("empty data", func(t *testing.T) {
		project := domain.NewProject("Dig a hole")
		projectRepo := in_memory.NewProjectRepository([]*domain.Project{project})
		insightRepo := in_memory.NewInsightRepository([]*domain.Insight{})
		uow := in_memory.InMemoryUnitOfWork{
			Projects: projectRepo,
			Insights: insightRepo,
		}
		apiHandler := NewInsightAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		url := fmt.Sprintf("/api/insights/?project_id=%s", project.ID)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("missing project id", func(t *testing.T) {
		insightRepo := in_memory.NewInsightRepository([]*domain.Insight{})
		uow := in_memory.InMemoryUnitOfWork{
			Insights: insightRepo,
		}
		apiHandler := NewInsightAPIHandler(&uow)

		mux := http.NewServeMux()
		apiHandler.Register(mux)

		req := httptest.NewRequest(http.MethodGet, "/api/insights/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
