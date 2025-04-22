package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type EntryHandler struct {
	entryService *service.EntryService
}

func NewEntryHandler(
	entryRepo domain.EntryRepository,
	projectRepo domain.ProjectReporitory,
) *EntryHandler {
	return &EntryHandler{
		entryService: service.NewEntryService(entryRepo, projectRepo),
	}
}

func (h *EntryHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/projects/{id}/entries", h.listEntries)
	mux.HandleFunc("POST /api/projects/{id}/entries", h.createEntry)
}

func (h *EntryHandler) listEntries(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	entries, err := h.entryService.ListEntries(id)
	if err != nil {
		http.Error(w, "failed to load entries", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}

func (h *EntryHandler) createEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var input service.CreateEntryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.entryService.CreateEntry(id, &input); err != nil {
		http.Error(w, "failed to create entry", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
