package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/db"
	"github.com/ScruffyPete/gologbook/internal/handlers"
)

func main() {
	fmt.Println("LOGBOOK!")

	mux := http.NewServeMux()
	h := &handlers.Handler{
		ProjectRepo: db.NewProjectRepository(),
	}
	handlers.RegisterProjectRoutes(mux, h)

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
