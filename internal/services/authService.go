package services

import (
	"context"

	"github.com/eif-courses/restaurant/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService struct {
	queries *repository.Queries
}

// paprasta funkcija (atitikmuo konstruktorius)
func NewAuthService(queries *repository.Queries) *AuthService {
	return &AuthService{queries: queries}
}

func (a *AuthService) CreateUser(ctx context.Context, username string, password string) (*repository.CreateUserRow, error) {

	params := repository.CreateUserParams{
		Email:    username,
		Password: password,
	}
	user, err := a.queries.CreateUser(ctx, params)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthService) CreateSession(ctx context.Context,
	userId int32,
	sessionToken string,
	expiresAt pgtype.Timestamp) (*repository.Session, error) {

	params := repository.CreateSessionParams{
		UserID:       userId,
		SessionToken: sessionToken,
		ExpiresAt:    expiresAt,
	}

	session, err := a.queries.CreateSession(ctx, params)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

//SessionToken string           `json:"session_token"`
//ExpiresAt    pgtype.Timestamp `json:"expires_at"`

func (a *AuthService) GetUserSessionByToken(ctx context.Context,
	sessionToken string,
	expiresAt pgtype.Timestamp) (*repository.GetSessionByTokenRow, error) {

	params := repository.GetSessionByTokenParams{
		SessionToken: sessionToken,
		ExpiresAt:    expiresAt,
	}

	session, err := a.queries.GetSessionByToken(ctx, params)

	if err != nil {
		return nil, err
	}
	return &session, nil
}
