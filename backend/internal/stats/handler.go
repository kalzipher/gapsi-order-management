package stats

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/http/response"
)

type Handler struct {
	service ServicePort
}

func NewHandler(service ServicePort) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Get(c *gin.Context) {
	filters := StatsFilters{
		Canal:           strings.TrimSpace(c.Query("canal")),
		Company:         strings.TrimSpace(c.Query("company")),
		FulfillmentType: strings.TrimSpace(c.Query("fulfillmentType")),
		ProductType:     strings.TrimSpace(c.Query("productType")),
	}

	result, err := h.service.Get(c.Request.Context(), filters)
	if err != nil {
		response.InternalServerError(c, "GET_STATS_ERROR", "could not get stats")
		return
	}

	response.OK(c, result)
}
