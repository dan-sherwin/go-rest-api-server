package restresponse

import "github.com/gin-gonic/gin"

func RestSuccess(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, data)
}

// RestSuccessNoContent sends a 204 No Content response.
func RestSuccessNoContent(c *gin.Context) {
	c.Status(204)
}
