package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	uow := in_memory.NewInMemoryUnitOfWork()
	apiHandler := NewAPIHandler(uow, nil)
	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	apiHandler.Register(mux)
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
