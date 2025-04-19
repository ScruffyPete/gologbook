package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/api"
	"github.com/ScruffyPete/gologbook/internal/db"
	"github.com/google/uuid"
)

type Handler struct {
	ProjectRepo db.ProjectReporitory
}

func RegisterProjectRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /projects", h.listProjects)
	mux.HandleFunc("GET /projects/{id}", h.getProjectById)
	mux.HandleFunc("POST /projects", h.createProjet)
	mux.HandleFunc("PATCH /projects/{id}", h.updateProjectDetails)
	mux.HandleFunc("DELETE /projects/{id}", h.deleteProject)
}

func (h *Handler) listProjects(w http.ResponseWriter, r *http.Request) {
	projects := h.ProjectRepo.ListProjects()
	json.NewEncoder(w).Encode(projects)
}

func (h *Handler) getProjectById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project := h.ProjectRepo.GetProject(id)
	if project == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(project)
}

func (h *Handler) createProjet(w http.ResponseWriter, r *http.Request) {
	var input api.ProjectRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	project := db.Project{
		ID:    uuid.NewString(),
		Title: input.Title,
	}

	if err := h.ProjectRepo.CreateProject(project); err != nil {
		http.Error(w, "failed to create project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateProjectDetails(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input api.ProjectRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	project := db.Project{
		ID:    id,
		Title: input.Title,
	} // FIXME
	if err := h.ProjectRepo.UpdateProject(project); err != nil {
		http.Error(w, "failed to update project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.ProjectRepo.DeleteProject(id); err != nil {
		http.Error(w, "failed to delete project", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
