package middlewares

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
)

// RequestLogger is a middleware for logging details of HTTP requests, including method, path, status, content length, referrer, user agent, and request body if applicable, excluding GET requests. It ensures detailed request monitoring and is designed for use with the Gin framework.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.Request.UserAgent()
		referrer := c.Request.Referer()
		method := c.Request.Method
		path := c.Request.URL.Path
		length := c.Request.ContentLength
		var postfield slog.Attr
		if c.Request.Header.Get("Content-Type") != "" && strings.HasPrefix(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
			postfield = slog.String("body", "multipart/form-data")
		} else {
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			bodyMap := make(map[string]interface{})
			err := json.Unmarshal(bodyBytes, &bodyMap)
			if err != nil {
				postfield = slog.String("body", string(bodyBytes))
			} else {
				postfield = slog.Any("body", bodyMap)
			}
		}
		c.Next()
		status := c.Writer.Status()
		if method != http.MethodGet {
			slog.Info("HTTP request",
				slog.String("method", method),
				slog.String("path", path),
				slog.Int("status", status),
				slog.Int64("length", length),
				slog.String("referrer", referrer),
				slog.String("useragent", userAgent),
				postfield,
			)
		}
	}
}
