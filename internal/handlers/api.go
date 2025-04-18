package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ScruffyPete/gologbook/api"
	"github.com/ScruffyPete/gologbook/internal/db"
)

func HandleProjectRoutes(mux *http.ServeMux, repo db.ProjectReporitory) {
	mux.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
		projects := repo.ListProjects()
		json.NewEncoder(w).Encode(projects)
	})

	mux.HandleFunc("GET /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		project := repo.GetProjectByID(id)
		if project == nil {
			http.Error(w, "proejct not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(project)
	})

	mux.HandleFunc("POST /projects", func(w http.ResponseWriter, r *http.Request) {
		var input api.CreateProjectRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		if err := repo.CreateProject(input.Title); err != nil {
			http.Error(w, "failed to create project", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := repo.DeleteProject(id); err != nil {
			http.Error(w, "failed to delete project", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("PATCH /projects/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := repo.LogEntry(id)
		if err != nil {
			http.Error(w, "failed to log entry", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}
