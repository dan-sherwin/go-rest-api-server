package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// epoch is a string representation of the Unix epoch time formatted using the HTTP Time Format standard in UTC.
var epoch = time.Unix(0, 0).UTC().Format(http.TimeFormat)

// noCacheHeaders defines a set of HTTP headers used to disable caching by preventing the storage and reuse of responses in browsers and intermediate caches.
var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, no-store, no-transform, must-revalidate, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

// etagHeaders defines a list of HTTP header names related to entity tags (ETag) and conditional requests for cache validation and control.
var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// NoCache is a middleware function that sets headers to prevent HTTP response caching and removes ETag-related request headers.
func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Delete any ETag headers that may have been set
		for _, v := range etagHeaders {
			if c.Request.Header.Get(v) != "" {
				c.Request.Header.Del(v)
			}
		}

		// Save our NoCache headers
		for k, v := range noCacheHeaders {
			c.Writer.Header().Set(k, v)
		}
		c.Next()
	}
}
