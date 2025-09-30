package handlers

import "net/http"

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return new(AuthHandler)
}

func (a *AuthHandler) GetSessions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
