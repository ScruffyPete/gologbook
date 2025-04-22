package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectRepo domain.ProjectReporitory) *ProjectHandler {
	return &ProjectHandler{
		projectService: service.NewProjectService(projectRepo),
	}
}

func (h *ProjectHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /projects", h.listProjects)
	mux.HandleFunc("GET /projects/{id}", h.getProjectById)
	mux.HandleFunc("POST /projects", h.createProjet)
	mux.HandleFunc("PATCH /projects/{id}", h.updateProjectDetails)
	mux.HandleFunc("DELETE /projects/{id}", h.deleteProject)
}

func (h *ProjectHandler) listProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.ListProjects()
	if err != nil {
		http.Error(w, "failed to list projects", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) getProjectById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project, err := h.projectService.GetProject(id)
	if err != nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) createProjet(w http.ResponseWriter, r *http.Request) {
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

func (h *ProjectHandler) updateProjectDetails(w http.ResponseWriter, r *http.Request) {
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

func (h *ProjectHandler) deleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.projectService.DeleteProject(id); err != nil {
		http.Error(w, "failed to delete project", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
