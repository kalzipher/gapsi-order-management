package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/http/middleware"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/http/response"
)

func (a *App) registerRoutes() {
	a.router.GET("/health", func(c *gin.Context) {
		response.OK(c, gin.H{
			"status": "ok",
		})
	})

	api := a.router.Group("/api")

	a.registerPublicModules(api)

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(
		NewAccessTokenValidatorAdapter(a.jwtService),
	))

	a.registerProtectedModules(protected)

	a.router.NoRoute(func(c *gin.Context) {
		response.Error(c, http.StatusNotFound, "ROUTE_NOT_FOUND", "route not found")
	})
}
