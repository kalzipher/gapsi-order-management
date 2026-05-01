package orders

import (
	"strconv"
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

func (h *Handler) List(c *gin.Context) {
	page, err := parsePositiveInt(c.DefaultQuery("page", "1"))
	if err != nil {
		response.BadRequest(c, "INVALID_PAGE", "page must be a positive integer")
		return
	}

	pageSize, err := parsePositiveInt(c.DefaultQuery("pageSize", "20"))
	if err != nil {
		response.BadRequest(c, "INVALID_PAGE_SIZE", "pageSize must be a positive integer")
		return
	}

	filters := ListOrdersFilters{
		Page:            page,
		PageSize:        pageSize,
		Canal:           strings.TrimSpace(c.Query("canal")),
		Company:         strings.TrimSpace(c.Query("company")),
		FulfillmentType: strings.TrimSpace(c.Query("fulfillmentType")),
		ProductType:     strings.TrimSpace(c.Query("productType")),
	}

	res, err := h.service.List(c.Request.Context(), filters)
	if err != nil {
		response.InternalServerError(c, "LIST_ORDERS_ERROR", "could not list orders")
		return
	}

	response.OK(c, res)
}

func (h *Handler) GetFilters(c *gin.Context) {
	result, err := h.service.GetFilters(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "GET_ORDER_FILTERS_ERROR", "could not get order filters")
		return
	}

	response.OK(c, result)
}

func parsePositiveInt(value string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	if parsed <= 0 {
		return 0, strconv.ErrSyntax
	}

	return parsed, nil
}
