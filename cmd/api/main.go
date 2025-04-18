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

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	repo, err := db.NewRepository()
	if err != nil {
		fmt.Println(err.Error())
	}

	if repo == nil {
		panic("Couln't setup the storage...")
	}

	handlers.HandleProjectRoutes(mux, repo)

	fmt.Println("Starting GoLogbook service...")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
