package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/db/postgres"
	"github.com/ScruffyPete/gologbook/internal/handler"
)

func main() {
	fmt.Println("LOGBOOK!")

	uow, err := postgres.NewPostgresUnitOfWork()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer uow.Close()

	projectHandler := handler.NewProjectAPIHandler(uow)
	entryHandler := handler.NewEntryAPIHandler(uow)

	mux := http.NewServeMux()
	projectHandler.Register(mux)
	entryHandler.Register(mux)

	wrappedMux := handler.JSONMiddleware(mux)

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, wrappedMux); err != nil {
		fmt.Println(err.Error())
	}
}
