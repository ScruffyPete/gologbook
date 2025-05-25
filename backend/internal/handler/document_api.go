package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type DocumentAPIHandler struct {
	documentService *service.DocumentService
}

func NewDocumentAPIHandler(uow domain.UnitOfWork) *DocumentAPIHandler {
	return &DocumentAPIHandler{
		documentService: service.NewDocumentService(uow),
	}
}

func (h *DocumentAPIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	wrappedMux := NewMiddlewareMux(middlewares...)

	wrappedMux.HandleFunc("GET /", h.listDocuments)

	mux.Handle("/api/documents/", http.StripPrefix("/api/documents", wrappedMux))
}

func (h *DocumentAPIHandler) listDocuments(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "project_id is required", http.StatusBadRequest)
		return
	}

	documents, err := h.documentService.ListDocuments(r.Context(), projectID)
	if err != nil {
		http.Error(w, "failed to load documents", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(documents)
}
