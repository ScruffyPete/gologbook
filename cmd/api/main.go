package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/handler"
)

func main() {
	fmt.Println("LOGBOOK!")

	projectRepo := in_memory.NewProjectRepository(nil)
	entryRepo := in_memory.NewEntryRepository(nil)
	projectHandler := handler.NewProjectHandler(projectRepo)
	entryHandler := handler.NewEntryHandler(entryRepo, projectRepo)

	mux := http.NewServeMux()
	projectHandler.Register(mux)
	entryHandler.Register(mux)

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
