package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ScruffyPete/gologbook/internal/db/in_memory"
	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestAuthAPIHandler_SignUp(t *testing.T) {
	t.Run("returns a 201 status code", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "test@example.com", "password": "password", "confirmPassword": "password"}`
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("returns a 400 status code if the password and confirm do not match", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "test@example.com", "password": "password", "confirmPassword": "password1"}`
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("returns a 400 status code if the email is not valid", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "test", "password": "password", "confirmPassword": "password"}`
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("returns a 500 status code if the user already exists", func(t *testing.T) {
		email := "test@example.com"
		password := "password"
		user := domain.NewUser(email, password)
		uow := in_memory.NewInMemoryUnitOfWork()
		uow.Users.CreateUser(user)
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "` + email + `", "password": "` + password + `", "confirmPassword": "` + password + `"}`
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestAuthAPIHandler_Login(t *testing.T) {
	t.Run("returns a 200 status code", func(t *testing.T) {
		email := "test@example.com"
		password := "password"
		hashedPassword, _ := service.HashPassword(password)
		user := domain.NewUser(email, hashedPassword)
		uow := in_memory.NewInMemoryUnitOfWork()
		uow.Users.CreateUser(user)
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "` + email + `", "password": "` + password + `"}`
		req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("returns a 401 status code if the user does not exist", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "test@example.com", "password": "password"}`
		req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("returns a 400 status code if the email is not valid", func(t *testing.T) {
		uow := in_memory.NewInMemoryUnitOfWork()
		apiHandler := NewAPIHandler(uow)
		mux := http.NewServeMux()
		apiHandler.Register(mux)

		payload := `{"email": "invalid-email", "password": "password"}`
		req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
