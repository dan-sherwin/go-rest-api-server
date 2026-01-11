package middlewares

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestLoggerOptions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		config        RequestLoggerConfig
		method        string
		expectedLog   bool
		expectedLevel slog.Level
	}{
		{
			name: "Default (no log GET, Info level)",
			config: RequestLoggerConfig{
				LogGetRequests: false,
				LogLevel:       LevelInfo,
			},
			method:        http.MethodGet,
			expectedLog:   false,
			expectedLevel: LevelInfo,
		},
		{
			name: "Log GET enabled, Info level",
			config: RequestLoggerConfig{
				LogGetRequests: true,
				LogLevel:       LevelInfo,
			},
			method:        http.MethodGet,
			expectedLog:   true,
			expectedLevel: LevelInfo,
		},
		{
			name: "Log POST (always logged), Debug level",
			config: RequestLoggerConfig{
				LogGetRequests: false,
				LogLevel:       LevelDebug,
			},
			method:        http.MethodPost,
			expectedLog:   true,
			expectedLevel: LevelDebug,
		},
		{
			name: "Log Off",
			config: RequestLoggerConfig{
				LogGetRequests: true,
				LogLevel:       LevelOff,
			},
			method:      http.MethodPost,
			expectedLog: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
			logger := slog.New(handler)
			slog.SetDefault(logger)

			r := gin.New()
			r.Use(RequestLoggerWithConfig(tt.config))
			r.Any("/test", func(c *gin.Context) {
				c.Status(200)
			})

			req, _ := http.NewRequest(tt.method, "/test", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if tt.expectedLog {
				assert.NotEmpty(t, buf.String())
				switch tt.expectedLevel {
				case LevelDebug:
					assert.Contains(t, buf.String(), "\"level\":\"DEBUG\"")
				case LevelInfo:
					assert.Contains(t, buf.String(), "\"level\":\"INFO\"")
				}
			} else {
				assert.Empty(t, buf.String())
			}
		})
	}

	t.Run("Global Config", func(t *testing.T) {
		// Save current globals
		oldLogGet := LogGetRequests
		oldLogLevel := LogLevel
		defer func() {
			LogGetRequests = oldLogGet
			LogLevel = oldLogLevel
		}()

		LogGetRequests = true
		LogLevel = LevelDebug

		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
		logger := slog.New(handler)
		slog.SetDefault(logger)

		r := gin.New()
		r.Use(RequestLogger())
		r.Any("/test", func(c *gin.Context) {
			c.Status(200)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.NotEmpty(t, buf.String())
		assert.Contains(t, buf.String(), "\"level\":\"DEBUG\"")
	})
}
