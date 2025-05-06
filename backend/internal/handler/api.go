package handler

import (
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
)

type APIHandler struct {
	authHandler    *AuthAPIHandler
	projectHandler *ProjectAPIHandler
	entryHandler   *EntryAPIHandler
}

func NewAPIHandler(uow domain.UnitOfWork, queue domain.Queue) *APIHandler {
	return &APIHandler{
		authHandler:    NewAuthAPIHandler(uow),
		projectHandler: NewProjectAPIHandler(uow),
		entryHandler:   NewEntryAPIHandler(uow, queue),
	}
}

func (h *APIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	h.authHandler.Register(mux)
	h.projectHandler.Register(mux, middlewares...)
	h.entryHandler.Register(mux, middlewares...)
}
