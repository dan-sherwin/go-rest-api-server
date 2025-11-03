// Package restresponse defines response codes and helpers for consistent JSON API responses.
package restresponse

import (
	"net/http"
	"strconv"
)

// Code represents an application-level response code used to map to HTTP status codes and string labels.
type Code uint32

const (
	// OK is returned on success.
	OK Code = iota

	// OkNoContent indicates that the request has succeeded and that there is no additional content to send in the response payload body.
	OkNoContent

	// Canceled indicates the operation was canceled (typically by the caller).
	// The gRPC framework will generate this error code when cancellation is requested.
	Canceled

	// Unknown error. See comments above for details.
	Unknown

	// InvalidArgument indicates client specified an invalid argument.
	InvalidArgument

	// DeadlineExceeded means operation expired before completion.
	DeadlineExceeded

	// NotFound means some requested entity (e.g., file or directory) was not found.
	NotFound

	// AlreadyExists means an attempt to create an entity failed because one already exists.
	AlreadyExists

	// PermissionDenied indicates the caller does not have permission to execute the specified operation.
	PermissionDenied

	// ResourceExhausted indicates some resource has been exhausted.
	ResourceExhausted

	// FailedPrecondition indicates operation was rejected because the system is not in a required state.
	FailedPrecondition

	// Aborted indicates the operation was aborted, typically due to a concurrency issue.
	Aborted

	// OutOfRange means operation was attempted past the valid range.
	OutOfRange

	// Unimplemented indicates operation is not implemented or not supported/enabled in this service.
	Unimplemented

	// Internal errors. Means some invariants expected by underlying system has been broken.
	Internal

	// Unavailable indicates the service is currently unavailable.
	Unavailable

	// DataLoss indicates unrecoverable data loss or corruption.
	DataLoss

	// Unauthenticated indicates the request does not have valid authentication credentials.
	Unauthenticated

	// BadRequest maps to HTTP 400 Bad Request.
	BadRequest
	// UnsupportedMediaType maps to HTTP 415 Unsupported Media Type.
	UnsupportedMediaType
	// NotAcceptable maps to HTTP 406 Not Acceptable.
	NotAcceptable
	// PayloadTooLarge maps to HTTP 413 Payload Too Large.
	PayloadTooLarge
	// TooManyRequests maps to HTTP 429 Too Many Requests.
	TooManyRequests
	// UnprocessableContent maps to HTTP 422 Unprocessable Content.
	UnprocessableContent
)

var httpStatusFromCode = map[Code]int{
	OK:                   http.StatusOK,
	OkNoContent:          http.StatusNoContent,
	Canceled:             499,
	Unknown:              http.StatusInternalServerError,
	InvalidArgument:      http.StatusBadRequest,
	DeadlineExceeded:     http.StatusGatewayTimeout,
	NotFound:             http.StatusNotFound,
	AlreadyExists:        http.StatusConflict,
	PermissionDenied:     http.StatusForbidden,
	Unauthenticated:      http.StatusUnauthorized,
	ResourceExhausted:    http.StatusTooManyRequests,
	FailedPrecondition:   http.StatusBadRequest, // intentionally not HTTP 412
	Aborted:              http.StatusConflict,
	OutOfRange:           http.StatusBadRequest,
	Unimplemented:        http.StatusNotImplemented,
	Internal:             http.StatusInternalServerError,
	Unavailable:          http.StatusServiceUnavailable,
	DataLoss:             http.StatusInternalServerError,
	BadRequest:           http.StatusBadRequest,
	UnsupportedMediaType: http.StatusUnsupportedMediaType,
	NotAcceptable:        http.StatusNotAcceptable,
	PayloadTooLarge:      http.StatusRequestEntityTooLarge,
	TooManyRequests:      http.StatusTooManyRequests,
	UnprocessableContent: http.StatusUnprocessableEntity,
}

var codeNames = []string{
	"OK",
	"OKNoContent",
	"Canceled",
	"Unknown",
	"InvalidArgument",
	"DeadlineExceeded",
	"NotFound",
	"AlreadyExists",
	"PermissionDenied",
	"ResourceExhausted",
	"FailedPrecondition",
	"Aborted",
	"OutOfRange",
	"Unimplemented",
	"Internal",
	"Unavailable",
	"DataLoss",
	"Unauthenticated",
	"BadRequest",
	"UnsupportedMediaType",
	"NotAcceptable",
	"PayloadTooLarge",
	"TooManyRequests",
	"UnprocessableContent",
}

// HTTPStatusFromCode converts a package Code into the corresponding HTTP status code.
func HTTPStatusFromCode(code Code) int {
	if s, ok := httpStatusFromCode[code]; ok {
		return s
	}
	return http.StatusInternalServerError
}

func (c Code) String() string {
	i := int(c)
	if i >= 0 && i < len(codeNames) {
		return codeNames[i]
	}
	return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
}
