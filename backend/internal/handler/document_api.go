package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type DocumentAPIHandler struct {
	documentService *service.DocumentService
}

func NewDocumentAPIHandler(uow domain.UnitOfWork, queue domain.Queue) *DocumentAPIHandler {
	return &DocumentAPIHandler{
		documentService: service.NewDocumentService(uow, queue),
	}
}

func (h *DocumentAPIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	wrappedMux := NewMiddlewareMux(middlewares...)

	wrappedMux.HandleFunc("GET /", h.listDocuments)
	wrappedMux.HandleFunc("GET /{id}/stream/", h.streamDocument)

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

func (h *DocumentAPIHandler) streamDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context()
	projectID := r.PathValue("id")
	if projectID == "" {
		http.Error(w, "project_id is required", http.StatusBadRequest)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	tokenStream := h.documentService.ConsumeDocumentStream(ctx, projectID)
	// fmt.Fprintf(w, "data: [[START]]\n\n")
	// flusher.Flush()

	for {
		select {
		case <-ctx.Done():
			return
		case token, ok := <-tokenStream:
			if !ok || token == "[[STOP]]" {
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", token)
			flusher.Flush()
		}
	}
}
