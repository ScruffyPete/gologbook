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

	wrappedMux.HandleFunc("GET /{projectID}/output/", h.listDocuments)
	wrappedMux.HandleFunc("GET /{projectID}/stream/", h.streamDocument)

	mux.Handle("/api/documents/", http.StripPrefix("/api/documents", wrappedMux))
}

func (h *DocumentAPIHandler) listDocuments(w http.ResponseWriter, r *http.Request) {
	projectID := r.PathValue("projectID")

	documents, err := h.documentService.GetLatestDocument(r.Context(), projectID)
	if err != nil {
		http.Error(w, "failed to load documents: "+err.Error(), http.StatusNotFound)
		// http.Error(w, "failed to load documents", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(documents)
}

func (h *DocumentAPIHandler) streamDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context()
	projectID := r.PathValue("projectID")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	tokenStream := h.documentService.ConsumeDocumentStream(ctx, projectID)

	for {
		select {
		case <-ctx.Done():
			return
		case token, ok := <-tokenStream:
			if !ok {
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", token)
			flusher.Flush()
		}
	}
}
