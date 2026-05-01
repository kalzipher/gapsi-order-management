package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/auth"
)

func (a *App) registerAuthModule(api *gin.RouterGroup) {
	authRepo := auth.NewRepository(a.db)

	jwtService := auth.NewJWTService(
		a.cfg.JWTAccessSecret,
		a.cfg.JWTRefreshSecret,
		a.cfg.JWTAccessTTL,
		a.cfg.JWTRefreshTTL,
	)

	authService := auth.NewService(authRepo, jwtService)

	a.jwtService = jwtService

	err := authService.EnsureAdminUser(
		context.Background(),
		a.cfg.AdminEmail,
		a.cfg.AdminPassword,
		a.cfg.AdminName,
	)
	if err != nil {
		panic(fmt.Errorf("could not ensure admin user: %w", err))
	}

	auth.NewHandler(authService).RegisterRoutes(api)
}
