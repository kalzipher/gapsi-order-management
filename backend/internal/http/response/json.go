package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorBody struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func JSON(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}

func Message(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, SuccessResponse{
		Message: message,
	})
}

func Error(c *gin.Context, statusCode int, code ErrorCode, message string) {
	c.JSON(statusCode, ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}

func BadRequest(c *gin.Context, code ErrorCode, message string) {
	Error(c, http.StatusBadRequest, code, message)
}

func Unauthorized(c *gin.Context, code ErrorCode, message string) {
	Error(c, http.StatusUnauthorized, code, message)
}

func Forbidden(c *gin.Context, code ErrorCode, message string) {
	Error(c, http.StatusForbidden, code, message)
}

func NotFound(c *gin.Context, code ErrorCode, message string) {
	Error(c, http.StatusNotFound, code, message)
}

func InternalServerError(c *gin.Context, code ErrorCode, message string) {
	Error(c, http.StatusInternalServerError, code, message)
}
