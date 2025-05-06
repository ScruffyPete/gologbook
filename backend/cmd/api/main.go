package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/db/postgres"
	"github.com/ScruffyPete/gologbook/internal/handler"
	"github.com/ScruffyPete/gologbook/internal/queue"
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
	queue := queue.NewInMemoryQueue()
	apiHandler := handler.NewAPIHandler(uow, queue)
	apiHandler.Register(mux, handler.AuthMiddleware)

	fmt.Println("Starting GoLogbook service...")

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
