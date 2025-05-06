package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type EntryAPIHandler struct {
	entryService *service.EntryService
}

func NewEntryAPIHandler(uow domain.UnitOfWork, queue domain.Queue) *EntryAPIHandler {
	return &EntryAPIHandler{
		entryService: service.NewEntryService(uow, queue),
	}
}

func (h *EntryAPIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	wrappedMux := NewMiddlewareMux(middlewares...)

	wrappedMux.HandleFunc("GET /", h.listEntries)
	wrappedMux.HandleFunc("POST /", h.createEntry)

	mux.Handle("/api/entries/", http.StripPrefix("/api/entries", wrappedMux))
}

func (h *EntryAPIHandler) listEntries(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "project_id is required", http.StatusBadRequest)
		return
	}

	entries, err := h.entryService.ListEntries(r.Context(), projectID)
	if err != nil {
		http.Error(w, "failed to load entries", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}

func (h *EntryAPIHandler) createEntry(w http.ResponseWriter, r *http.Request) {
	var input service.CreateEntryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	entry, err := h.entryService.CreateEntry(r.Context(), &input)
	if err != nil {
		http.Error(w, "failed to create entry", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}
