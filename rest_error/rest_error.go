package rest_error

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
	RestErrorRespond(c, Internal, message, details)
}

func RestSuccessRespond(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}
