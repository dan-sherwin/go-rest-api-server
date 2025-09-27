package rest_response

import (
	"github.com/gin-gonic/gin"
)

func RestErrorRespond(c *gin.Context, code Code, message string, details ...interface{}) {
	c.Header("Content-Type", "application/json")
	httpStatus := HTTPStatusFromCode(code)
	c.JSON(httpStatus, map[string]any{
		"code":    code.String(),
		"message": message,
		"details": details,
	})
}

func RestUnknownErrorRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, Internal, message, details...)
}

// Convenience helpers for new HTTP-aligned error codes
func RestBadRequestRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, BadRequest, message, details...)
}

func RestUnsupportedMediaTypeRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, UnsupportedMediaType, message, details...)
}

func RestNotAcceptableRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, NotAcceptable, message, details...)
}

func RestPayloadTooLargeRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, PayloadTooLarge, message, details...)
}

func RestTooManyRequestsRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, TooManyRequests, message, details...)
}

func RestUnprocessableContentRespond(c *gin.Context, message string, details ...interface{}) {
	RestErrorRespond(c, UnprocessableContent, message, details...)
}
