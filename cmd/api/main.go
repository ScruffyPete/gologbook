package main

import (
	"fmt"
	"net/http"

	"github.com/ScruffyPete/gologbook/internal/handlers"
)

func main() {
	fmt.Println("LOGBOOK!")

	var mux *http.ServeMux = http.NewServeMux()

	handlers.Handler(mux)

	fmt.Println("Starting GoLogbook service...")

	err := http.ListenAndServe("localhost:8000", mux)

	if err != nil {
		fmt.Println(err.Error())
	}
}
