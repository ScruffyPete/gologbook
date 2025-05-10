package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/db/postgres"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/handler"
	"github.com/ScruffyPete/gologbook/internal/queue"
)

func main() {
	fmt.Println("LOGBOOK!")
	fmt.Println("Starting service...")

	inMemory := flag.Bool("in-memory", false, "Use in-memory storage")
	flag.Parse()

	var uow domain.UnitOfWork
	var q domain.Queue
	var err error
	if *inMemory {
		fmt.Println("Using in-memory storage")
		uow = in_memory.NewInMemoryUnitOfWork()
		q = queue.NewInMemoryQueue()
	} else {
		fmt.Println("Using postgres storage")
		uow, err = postgres.NewPostgresUnitOfWork()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Using redis queue")
		q, err = queue.NewRedisQueue()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	defer uow.Close()
	defer q.Close()

	mux := http.NewServeMux()
	apiHandler := handler.NewAPIHandler(uow, q)
	apiHandler.Register(mux, handler.AuthMiddleware)

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
