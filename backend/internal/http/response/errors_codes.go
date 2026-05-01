package response

type ErrorCode string

const ( // General errors
	ErrInternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrBadRequest          ErrorCode = "BAD_REQUEST"
	ErrUnauthorized        ErrorCode = "UNAUTHORIZED"
	ErrForbidden           ErrorCode = "FORBIDDEN"
	ErrNotFound            ErrorCode = "NOT_FOUND"
)

//INVALID_REQUEST
