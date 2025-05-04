package handler

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/ScruffyPete/gologbook/internal/domain"
	"github.com/ScruffyPete/gologbook/internal/service"
)

type AuthAPIHandler struct {
	authService *service.AuthService
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthAPIHandler(uow domain.UnitOfWork) *AuthAPIHandler {
	if uow == nil {
		panic("uow is nil")
	}
	return &AuthAPIHandler{authService: service.NewAuthService(uow)}
}

func (h *AuthAPIHandler) signUp(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(input.Email); err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	if input.Password != input.Confirm {
		http.Error(w, "password and confirm do not match", http.StatusBadRequest)
		return
	}

	err := h.authService.SignUp(r.Context(), input.Email, input.Password)
	if err != nil {
		http.Error(w, "failed to register", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AuthAPIHandler) login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(input.Email); err != nil {
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), input.Email, input.Password)
	if err != nil {
		http.Error(w, "failed to login", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
