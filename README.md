# Go Rest API Server

A lightweight helper package for building HTTP/HTTPS services with [Gin](https://github.com/gin-gonic/gin). It provides sensible defaults, production‑friendly middleware, convenient response helpers, and support for listening on multiple interfaces and/or ports. It preserves backward compatibility with legacy single‑address APIs.

- Multi‑address listeners for HTTP and HTTPS
- Optional self‑signed TLS certificate generation for development
- Minimal, composable middlewares: CORS, structured request logging, no‑cache headers
- Small convenience helpers for consistent JSON responses
- Tested, linted, and vulnerability‑checked in CI

## Installation

Requirements: Go 1.25.3 or newer.

```
go get github.com/dan-sherwin/go-rest-api-server@v0.4.1
```

## Quick Start

### Single address (HTTP)
```go
package main

import (
    "log"
    restapi "github.com/dan-sherwin/go-rest-api-server"
)

func main() {
    restapi.ListeningAddress = "0.0.0.0:5555"
    restapi.StartHttpServer()
    defer func() {
        if err := restapi.ShutdownHttpServer(); err != nil {
            log.Fatal(err)
        }
    }()

    select {} // serve forever
}
```

### Multiple addresses (HTTP)
```go
restapi.ListeningAddresses = []string{"0.0.0.0:8080", "127.0.0.1:9090"}
restapi.StartHttpServer()
```

### HTTPS with self‑signed certificates (development)
```go
// Start HTTPS on multiple addresses and auto‑generate a self‑signed cert/key.
restapi.HTTPSListeningAddresses = []string{"0.0.0.0:8443", "127.0.0.1:9443"}
restapi.StartHttpsServer("", "", true)
// Optionally ensure cleanup if your app may exit without calling Shutdown.
defer restapi.CleanupGeneratedTLSFiles()
```

## API Overview

The package exposes a small set of globals and functions. Legacy names are preserved intentionally for backward compatibility.

- Configuration
  - `ListeningAddress string` — default `"0.0.0.0:5555"`
  - `HTTPSListeningAddress string` — default `"0.0.0.0:5556"`
  - `ListeningAddresses []string` — if non‑empty, takes precedence over `ListeningAddress`
  - `HTTPSListeningAddresses []string` — if non‑empty, takes precedence over `HTTPSListeningAddress`
- Server lifecycle
  - `StartHttpServer()` — start one or more HTTP servers
  - `StartHttpsServer(certFile, keyFile string, autoSelfSigned bool)` — start one or more HTTPS servers; when `autoSelfSigned` is true and cert/key are empty or missing, a dev self‑signed pair is generated and reused for all listeners
  - `ShutdownHttpServer() error` — graceful shutdown of all HTTP servers
  - `ShutdownHttpsServer() error` — graceful shutdown of all HTTPS servers and cleanup of any auto‑generated TLS files
  - `CleanupGeneratedTLSFiles()` — force cleanup of any auto‑generated TLS artifacts
- Routing and utilities
  - `RegisterRouters(funcs ...func(r *gin.Engine))` — add routes to the underlying Gin engine
  - `DisablePing()` — disables the default `/ping` route
  - `ClientIP(c *gin.Context) string` — returns `X-Forwarded-For` if present, otherwise `c.ClientIP()`

## Middleware

In `github.com/dan-sherwin/go-rest-api-server/middlewares`:

- `CORSMiddleware()` — permissive CORS headers with preflight handling
- `RequestLogger()` — structured request logging (method, path, status, content length, referrer, UA, and body for non‑GET requests)
- `NoCache()` — sets headers to prevent client/proxy caching

Use with your own Gin engine, or simply rely on the engine created by this package.

## Response Helpers

In `github.com/dan-sherwin/go-rest-api-server/restresponse`:

- `RestSuccess(c *gin.Context, data any)` — 200 JSON response
- `RestSuccessNoContent(c *gin.Context)` — 204 response
- Error helpers (map to sensible HTTP statuses):
  - `RestErrorRespond(c, code, message, details...)`
  - `RestUnknownErrorRespond`
  - `RestBadRequestRespond`
  - `RestUnsupportedMediaTypeRespond`
  - `RestNotAcceptableRespond`
  - `RestPayloadTooLargeRespond`
  - `RestTooManyRequestsRespond`
  - `RestUnprocessableContentRespond`

Status mapping is implemented via `restresponse.HTTPStatusFromCode` and `restresponse.Code` values.

## Development

- Run tests
  ```
  go test ./... -v
  ```
- Lint
  ```
  golangci-lint run --timeout 5m
  ```
- Security
  ```
  govulncheck ./...
  ```

## Versioning & Changelog

This project follows Semantic Versioning. See `CHANGELOG.md` for release notes.

## License

This project is licensed under the terms of the license found in the `LICENSE` file.
