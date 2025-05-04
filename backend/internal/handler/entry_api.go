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

func NewEntryAPIHandler(uow domain.UnitOfWork) *EntryAPIHandler {
	if uow == nil {
		panic("ProjectAPIHandler: unit of work cannot be nil")
	}
	return &EntryAPIHandler{
		entryService: service.NewEntryService(uow),
	}
}

func (h *EntryAPIHandler) listEntries(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	entries, err := h.entryService.ListEntries(r.Context(), id)
	if err != nil {
		http.Error(w, "failed to load entries", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(entries)
}

func (h *EntryAPIHandler) createEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var input service.CreateEntryInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	entry, err := h.entryService.CreateEntry(r.Context(), id, &input)
	if err != nil {
		http.Error(w, "failed to create entry", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}
