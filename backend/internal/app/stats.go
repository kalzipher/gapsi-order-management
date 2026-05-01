package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/stats"
)

func (a *App) registerStatsModule(api *gin.RouterGroup) {
	statsRepo := stats.NewRepository(a.db)
	statsService := stats.NewService(statsRepo)
	statsHandler := stats.NewHandler(statsService)

	api.Handle(http.MethodGet, "/stats", statsHandler.Get)
}
