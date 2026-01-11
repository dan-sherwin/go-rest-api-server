package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// LevelDebug corresponds to slog.LevelDebug.
	LevelDebug = slog.LevelDebug
	// LevelInfo corresponds to slog.LevelInfo.
	LevelInfo = slog.LevelInfo
	// LevelWarn corresponds to slog.LevelWarn.
	LevelWarn = slog.LevelWarn
	// LevelError corresponds to slog.LevelError.
	LevelError = slog.LevelError
	// LevelOff is a custom level used to disable request logging.
	LevelOff = slog.Level(100)
)

var (
	// LogGetRequests determines whether or not to log GET requests.
	LogGetRequests = false
	// LogLevel specifies the level at which requests should be logged.
	LogLevel = LevelInfo
)

// RequestLoggerConfig defines the configuration for the RequestLogger middleware.
type RequestLoggerConfig struct {
	// LogGetRequests determines whether or not to log GET requests.
	LogGetRequests bool
	// LogLevel specifies the level at which requests should be logged.
	LogLevel slog.Level
}

// RequestLogger is a middleware for logging details of HTTP requests using global configuration.
func RequestLogger() gin.HandlerFunc {
	return RequestLoggerWithDynamicConfig(&LogGetRequests, &LogLevel)
}

// RequestLoggerWithDynamicConfig is a middleware for logging details of HTTP requests with dynamic configuration.
func RequestLoggerWithDynamicConfig(logGetRequests *bool, logLevel *slog.Level) gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestLoggerWithConfig(RequestLoggerConfig{
			LogGetRequests: *logGetRequests,
			LogLevel:       *logLevel,
		})(c)
	}
}

// RequestLoggerWithConfig is a middleware for logging details of HTTP requests with custom configuration.
func RequestLoggerWithConfig(config RequestLoggerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.LogLevel == LevelOff {
			c.Next()
			return
		}
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
				bodyBytes, _ = io.ReadAll(c.Request.Body)
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
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
		if method != http.MethodGet || config.LogGetRequests {
			slog.Log(c.Request.Context(), config.LogLevel, "HTTP request",
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
