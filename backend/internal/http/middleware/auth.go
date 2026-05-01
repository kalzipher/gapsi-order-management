package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/http/response"
)

const (
	ContextUserIDKey = "userId"
	ContextEmailKey  = "email"
)

type AccessTokenClaims struct {
	UserID string
	Email  string
}

type AccessTokenValidator interface {
	ValidateAccessToken(token string) (*AccessTokenClaims, error)
}

func AuthMiddleware(tokenValidator AccessTokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(
				c,
				"MISSING_AUTHORIZATION_HEADER",
				"authorization header is required",
			)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Unauthorized(
				c,
				"INVALID_AUTHORIZATION_HEADER",
				"authorization header must use Bearer token",
			)
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(parts[1])
		if tokenString == "" {
			response.Unauthorized(
				c,
				"EMPTY_ACCESS_TOKEN",
				"access token is required",
			)
			c.Abort()
			return
		}

		claims, err := tokenValidator.ValidateAccessToken(tokenString)
		if err != nil {
			response.Unauthorized(
				c,
				"INVALID_ACCESS_TOKEN",
				"invalid or expired access token",
			)
			c.Abort()
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextEmailKey, claims.Email)

		c.Next()
	}
}
