package services

import "github.com/eif-courses/restaurant/internal/repository"

type AuthService struct {
	queries *repository.Queries
}

// paprasta funkcija (atitikmuo konstruktorius)
func NewAuthService(queries *repository.Queries) *AuthService {
	return &AuthService{queries: queries}
}
