package rest_api_server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/dan-sherwin/go-rest-api-server/middlewares"
	"github.com/gin-gonic/gin"
)

// httpApp is the instance of the Gin Engine used for HTTP routing and handling.
//
// httpServer represents the HTTP server that will serve requests over the network.
//
// ListeningAddress is the address and port the server listens on for incoming HTTP connections.
var (
	httpApp          *gin.Engine
	httpServer       *http.Server
	ListeningAddress = "0.0.0.0:5555"
	disablePing      = false
)

// init configures the HTTP server with necessary middlewares including CORS, request ID generation, logging, and no-cache headers, and initializes the Gin engine with recovery and trusted proxy settings.
func init() {
	slog.Debug("initializing http server")
	gin.SetMode(gin.ReleaseMode)
	httpApp = gin.New()
	httpApp.Use(gin.Recovery())
	_ = httpApp.SetTrustedProxies(nil)
	httpApp.Use(gin.Recovery(), middlewares.CORSMiddleware(), middlewares.RequestLogger(), middlewares.NoCache())
}

// StartHttpServer initializes and starts an HTTP server at the specified listening address.
// It sets up routes, logs server start information, and handles any errors during server startup.
// Runs the server in a separate goroutine to allow non-blocking operation.
func StartHttpServer() {
	slog.Info(fmt.Sprintf("starting service on http://%s", ListeningAddress))
	setupRoutes()
	httpServer = &http.Server{Addr: ListeningAddress, Handler: httpApp.Handler()}
	go func() {
		slog.Info("listening and serving HTTP on " + ListeningAddress)
		if err := httpServer.ListenAndServe(); err != nil {
			if err.Error() != "http: Server closed" {
				slog.Error("error occurred while setting up the server. " + err.Error())
				os.Exit(1)
			}
		}
	}()
}

// StartHttpTLSServer initializes and starts an HTTPS server with TLS using the provided certificate and key files.
// It listens on the globally defined ListeningAddress and handles HTTP routes through the configured httpApp router.
// The server runs asynchronously and logs relevant startup and error messages.
func StartHttpTLSServer(certFile, keyFile string) {
	slog.Info(fmt.Sprintf("starting tls http service on http://%s", ListeningAddress))
	setupRoutes()
	httpServer = &http.Server{Addr: ListeningAddress, Handler: httpApp.Handler()}
	go func() {
		slog.Info("listening and serving TLS HTTP on " + ListeningAddress)
		if err := httpServer.ListenAndServeTLS(certFile, keyFile); err != nil {
			if err.Error() != "http: Server closed" {
				slog.Error("error occurred while setting up the server. " + err.Error())
				os.Exit(1)
			}
		}
	}()
}

// DisablePing sets the internal `disablePing` variable to true, effectively disabling the ping functionality.
func DisablePing() {
	disablePing = true
}

// RegisterRouters registers multiple router configurations to the HTTP application engine. It accepts a variadic number of functions, each receiving the *gin.Engine instance for adding routes and middleware.
func RegisterRouters(registrars ...func(r *gin.Engine)) {
	for _, r := range registrars {
		r(httpApp)
	}
}

// ShutdownHttpServer gracefully shuts down the HTTP server, if it is running, using a background context. It returns an error if the shutdown process encounters any issues or if the server cannot be properly terminated.
func ShutdownHttpServer() error {
	if httpServer == nil {
		return nil
	}
	return httpServer.Shutdown(context.Background())
}

// setupRoutes defines HTTP routes and their corresponding handlers for the application.
func setupRoutes() {
	httpApp.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	})
	if !disablePing {
		httpApp.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}

// ClientIP retrieves the client IP address from the request context. It checks the "X-Forwarded-For" header first and returns its value if present; otherwise, it falls back to using c.ClientIP().
func ClientIP(c *gin.Context) string {
	buf := c.GetHeader("X-Forwarded-For")
	if buf != "" {
		return buf
	}
	return c.ClientIP()
}
