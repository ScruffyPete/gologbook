package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type InsightAPIHandler struct {
	insightService *service.InsightService
}

func NewInsightAPIHandler(uow domain.UnitOfWork) *InsightAPIHandler {
	return &InsightAPIHandler{
		insightService: service.NewInsightService(uow),
	}
}

func (h *InsightAPIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	wrappedMux := NewMiddlewareMux(middlewares...)

	wrappedMux.HandleFunc("GET /", h.listInsights)

	mux.Handle("/api/insights/", http.StripPrefix("/api/insights", wrappedMux))
}

func (h *InsightAPIHandler) listInsights(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "project_id is required", http.StatusBadRequest)
		return
	}

	insights, err := h.insightService.ListInsights(r.Context(), projectID)
	if err != nil {
		http.Error(w, "failed to load insights", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(insights)
}
