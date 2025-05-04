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

	mux := http.NewServeMux()
	apiHandler := handler.NewAPIHandler(uow)
	apiHandler.Register(mux, handler.AuthMiddleware)

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
