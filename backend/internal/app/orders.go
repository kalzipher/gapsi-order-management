package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/orders"
)

func (a *App) registerOrdersModule(api *gin.RouterGroup) {
	ordersRepo := orders.NewRepository(a.db)
	ordersService := orders.NewService(ordersRepo)
	ordersHandler := orders.NewHandler(ordersService)

	api.Handle(http.MethodGet, "/orders", ordersHandler.List)
	api.Handle(http.MethodGet, "/orders/filters", ordersHandler.GetFilters)
}
