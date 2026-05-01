package auth

import (
	"errors"
	"net/http"

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

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.POST("/login", h.Login)
	auth.POST("/refresh", h.Refresh)
	auth.POST("/logout", h.Logout)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			response.ErrBadRequest,
			"email and password are required",
		)
		return
	}

	res, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			response.Unauthorized(
				c,
				response.ErrUnauthorized,
				"invalid email or password",
			)
			return
		}

		response.InternalServerError(
			c,
			response.ErrInternalServerError,
			"could not login",
		)
		return
	}

	response.OK(c, res)
}

func (h *Handler) Refresh(c *gin.Context) {
	var req RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			response.ErrBadRequest,
			"refreshToken is required",
		)
		return
	}

	res, err := h.service.Refresh(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, ErrInvalidRefreshToken) {
			response.Unauthorized(
				c,
				response.ErrUnauthorized,
				"invalid or expired refresh token",
			)
			return
		}

		response.InternalServerError(
			c,
			response.ErrInternalServerError,
			"could not refresh token",
		)
		return
	}

	response.OK(c, res)
}

func (h *Handler) Logout(c *gin.Context) {
	var req LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			response.ErrBadRequest,
			"refreshToken is required",
		)
		return
	}

	if err := h.service.Logout(c.Request.Context(), req); err != nil {
		response.InternalServerError(
			c,
			response.ErrInternalServerError,
			"could not logout",
		)
		return
	}

	response.Message(c, http.StatusOK, "logout successful")
}
