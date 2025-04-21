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

	repo := in_memory.NewProjectRepository(nil)
	h := handler.NewHandler(repo)
	mux := h.NewRouter()

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
