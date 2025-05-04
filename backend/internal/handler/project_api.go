package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type ProjectAPIHandler struct {
	projectService *service.ProjectService
}

func NewProjectAPIHandler(uow domain.UnitOfWork) *ProjectAPIHandler {
	if uow == nil {
		panic("ProjectAPIHandler: unit of work cannot be nil")
	}

	return &ProjectAPIHandler{
		projectService: service.NewProjectService(uow),
	}
}

func (h *ProjectAPIHandler) listProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projectService.ListProjects(r.Context())
	if err != nil {
		http.Error(w, "failed to list projects", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectAPIHandler) getProjectById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project, err := h.projectService.GetProject(r.Context(), id)
	if err != nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(project)
}

func (h *ProjectAPIHandler) createProjet(w http.ResponseWriter, r *http.Request) {
	var input service.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	project, err := h.projectService.CreateProject(r.Context(), &input)
	if err != nil {
		http.Error(w, "failed to create project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectAPIHandler) updateProjectDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input service.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.projectService.UpdateProject(r.Context(), id, &input); err != nil {
		http.Error(w, "failed to update project", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProjectAPIHandler) deleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.projectService.DeleteProject(r.Context(), id); err != nil {
		http.Error(w, "failed to delete project", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
