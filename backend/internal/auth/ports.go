package auth

import (
	"context"
	"time"
)

type ServicePort interface {
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	Refresh(ctx context.Context, req RefreshRequest) (*RefreshResponse, error)
	Logout(ctx context.Context, req LogoutRequest) error
	EnsureAdminUser(ctx context.Context, email, password, name string) error
}

type RepositoryPort interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user User) error
	SaveRefreshToken(ctx context.Context, token RefreshToken) error
	FindValidRefreshToken(ctx context.Context, tokenHash string) (*RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, tokenHash string) error
	RevokeAllUserRefreshTokens(ctx context.Context, userID string) error
}

type TokenPort interface {
	GenerateAccessToken(user User) (string, time.Time, error)
	GenerateRefreshToken(user User) (string, time.Time, error)
	ValidateAccessToken(token string) (*Claims, error)
	ValidateRefreshToken(token string) (*Claims, error)
}
