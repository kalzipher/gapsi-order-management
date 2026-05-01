package app

import (
	"github.com/github/gapsi-order-management-dashboard/backend/internal/auth"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/http/middleware"
)

type AccessTokenValidatorAdapter struct {
	jwtService *auth.JWTService
}

func NewAccessTokenValidatorAdapter(jwtService *auth.JWTService) AccessTokenValidatorAdapter {
	return AccessTokenValidatorAdapter{
		jwtService: jwtService,
	}
}

func (a AccessTokenValidatorAdapter) ValidateAccessToken(token string) (*middleware.AccessTokenClaims, error) {
	claims, err := a.jwtService.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}

	return &middleware.AccessTokenClaims{
		UserID: claims.UserID,
		Email:  claims.Email,
	}, nil
}
