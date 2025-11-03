package restresponse

import (
	"github.com/gin-gonic/gin"
)

// RestErrorRespond sends a JSON error response with the given Code, message and optional details.
// The HTTP status is derived from the provided Code via HTTPStatusFromCode.
func RestErrorRespond(c *gin.Context, code Code, message string, details ...interface{}) {
	c.Header("Content-Type", "application/json")
	httpStatus := HTTPStatusFromCode(code)
	c.JSON(httpStatus, map[string]any{
		"code":    code.String(),
		"message": message,
		"details": details,
	})
}

// RestUnknownErrorRespond sends a generic 500-style error using the Internal code.
func RestUnknownErrorRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, Internal, message, details...)
}

// RestBadRequestRespond sends an HTTP 400 Bad Request style error.
func RestBadRequestRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, BadRequest, message, details...)
}

// RestUnsupportedMediaTypeRespond sends an HTTP 415 Unsupported Media Type style error.
func RestUnsupportedMediaTypeRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, UnsupportedMediaType, message, details...)
}

// RestNotAcceptableRespond sends an HTTP 406 Not Acceptable style error.
func RestNotAcceptableRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, NotAcceptable, message, details...)
}

// RestPayloadTooLargeRespond sends an HTTP 413 Payload Too Large style error.
func RestPayloadTooLargeRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, PayloadTooLarge, message, details...)
}

// RestTooManyRequestsRespond sends an HTTP 429 Too Many Requests style error.
func RestTooManyRequestsRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, TooManyRequests, message, details...)
}

// RestUnprocessableContentRespond sends an HTTP 422 Unprocessable Content style error.
func RestUnprocessableContentRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, UnprocessableContent, message, details...)
}
