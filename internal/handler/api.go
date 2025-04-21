package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type Handler struct {
	projectService *service.ProjectService
}

func NewHandler(repo domain.ProjectReporitory) *Handler {
	return &Handler{
		projectService: service.NewProjectService(repo),
	}
}

func (h *Handler) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /projects", h.listProjects)
	mux.HandleFunc("GET /projects/{id}", h.getProjectById)
	mux.HandleFunc("POST /projects", h.createProjet)
	mux.HandleFunc("PATCH /projects/{id}", h.updateProjectDetails)
	mux.HandleFunc("DELETE /projects/{id}", h.deleteProject)
	return mux
}

func (h *Handler) listProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.ListProjects()
	if err != nil {
		http.Error(w, "failed to list projects", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}

func (h *Handler) getProjectById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project, err := h.projectService.GetProject(id)
	if err != nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func (h *Handler) createProjet(w http.ResponseWriter, r *http.Request) {
	var input service.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.projectService.CreateProject(&input); err != nil {
		http.Error(w, "failed to create project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateProjectDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input service.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.projectService.UpdateProject(id, &input); err != nil {
		http.Error(w, "failed to update project", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.projectService.DeleteProject(id); err != nil {
		http.Error(w, "failed to delete project", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
