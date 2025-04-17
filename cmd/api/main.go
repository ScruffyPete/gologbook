package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ScruffyPete/gologbook/internal/handlers"
)

func main() {
	fmt.Println("LOGBOOK!")

	port := os.Getenv("PORT")

	var mux *http.ServeMux = http.NewServeMux()

	handlers.Handler(mux)

	fmt.Println("Starting GoLogbook service...")

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
