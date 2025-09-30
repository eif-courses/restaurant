package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eif-courses/restaurant/internal/services"
)

type AuthHandler struct {
	authServive *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (a *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	type CreateUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// .. check params is empty validuoti duomenis

	user, err := a.authServive.CreateUser(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(user)

}
