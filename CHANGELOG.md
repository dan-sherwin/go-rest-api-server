# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog and this project adheres to Semantic Versioning.

## [Unreleased]

- Nothing yet.

## [v0.4.1] - 2025-11-03

### Added
- Support for listening on multiple interface addresses and/or ports for both HTTP and HTTPS while preserving legacy single‑address API.
- CI‑friendly test suite covering:
  - HTTP server on a single legacy address (`/ping`).
  - HTTP server on multiple addresses (`/ping`).
  - HTTPS server with auto self‑signed certificates (`/ping`) and cleanup of generated files.
  - `ClientIP` behavior preferring `X-Forwarded-For`.

### Changed
- Internal refactors to support multiple server instances while keeping original global handles for backward compatibility.
- Documentation comments across packages for exported identifiers.

### Fixed
- Lint issues reported by `golangci-lint` (errcheck, revive, staticcheck, unused) without breaking API.

### Security
- Upgraded toolchain to Go 1.25.3 (via `go` directive and `toolchain`) to address standard library vulnerabilities (GO-2025-4007, GO-2025-4008, GO-2025-4009, GO-2025-4010, GO-2025-4011, GO-2025-4013). `govulncheck` reports no vulnerabilities.

[Unreleased]: https://github.com/dan-sherwin/go-rest-api-server/compare/v0.4.1...HEAD
[v0.4.1]: https://github.com/dan-sherwin/go-rest-api-server/releases/tag/v0.4.1
