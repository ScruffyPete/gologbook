package handlers

import (
	"fmt"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "return all projects")
}

func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "return project with id: %s", id)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new project")
}

func LogEntry(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "log a new entry to project with id: %s", id)
}
