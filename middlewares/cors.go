package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware sets Cross-Origin Resource Sharing (CORS) headers to allow requests across origins.
// Handles preflight requests by responding with status 200 for OPTIONS method.
// Configures allowed origins, methods, headers, credentials, and max age for caching.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowHeaders := c.Request.Header.Get("Access-Control-Request-Headers")
		if allowHeaders == "" {
			allowHeaders = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "600")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
