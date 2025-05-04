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

func NewAPIHandler(uow domain.UnitOfWork) *APIHandler {
	return &APIHandler{
		authHandler:    NewAuthAPIHandler(uow),
		projectHandler: NewProjectAPIHandler(uow),
		entryHandler:   NewEntryAPIHandler(uow),
	}
}

func (h *APIHandler) Register(mux *http.ServeMux, middlewares ...func(http.Handler) http.Handler) {
	mux.HandleFunc("POST /api/auth/signup", h.authHandler.signUp)
	mux.HandleFunc("POST /api/auth/login", h.authHandler.login)

	wrappedMux := NewMiddlewareMux(middlewares...)

	wrappedMux.HandleFunc("GET /projects", h.projectHandler.listProjects)
	wrappedMux.HandleFunc("GET /projects/{id}", h.projectHandler.getProjectById)
	wrappedMux.HandleFunc("POST /projects", h.projectHandler.createProjet)
	wrappedMux.HandleFunc("PATCH /projects/{id}", h.projectHandler.updateProjectDetails)
	wrappedMux.HandleFunc("DELETE /projects/{id}", h.projectHandler.deleteProject)

	wrappedMux.HandleFunc("GET /projects/{id}/entries", h.entryHandler.listEntries)
	wrappedMux.HandleFunc("POST /projects/{id}/entries", h.entryHandler.createEntry)

	mux.Handle("/api/", http.StripPrefix("/api", wrappedMux))
}
