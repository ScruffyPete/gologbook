package handlers

import (
	"net/http"
)

func Handler(mux *http.ServeMux) {
	mux.HandleFunc("GET /projects", GetProjects)
	mux.HandleFunc("GET /projects/{id}", GetProjectByID)
	mux.HandleFunc("POST /projects", CreateProject)
	mux.HandleFunc("PATCH /projects/{id}", LogEntry)
}
