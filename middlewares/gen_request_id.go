package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GenRequestID generates a middleware function that assigns a unique request ID to each HTTP request.
// The generated ID is stored in the request context, added to the request and response headers under "X-Spacelink-Request-ID",
// and ensures traceability of requests across systems.
func GenRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Request.Header.Set("X-Spacelink-Request-ID", id)
		c.Set("X-Spacelink-Request-ID", id)
		c.Header("X-Spacelink-Request-ID", id)
		c.Next()
	}
}
