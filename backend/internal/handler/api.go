package handler

import (
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type APIHandler struct {
	authHandler     *AuthAPIHandler
	projectHandler  *ProjectAPIHandler
	entryHandler    *EntryAPIHandler
	documentHandler *DocumentAPIHandler
}

func NewAPIHandler(uow domain.UnitOfWork, queue domain.Queue) *APIHandler {
	return &APIHandler{
		authHandler:     NewAuthAPIHandler(uow),
		projectHandler:  NewProjectAPIHandler(uow),
		entryHandler:    NewEntryAPIHandler(uow, queue),
		documentHandler: NewDocumentAPIHandler(uow, queue),
	}
}

func (h *APIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	mux.HandleFunc("GET /healthz", h.healthCheckHandler)
	h.authHandler.Register(mux)
	h.projectHandler.Register(mux, middlewares...)
	h.entryHandler.Register(mux, middlewares...)
	h.documentHandler.Register(mux, middlewares...)
}

func (h *APIHandler) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
