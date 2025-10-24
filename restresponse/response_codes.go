package restresponse

import (
	"net/http"
	"strconv"
)

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

	// Additional HTTP-aligned codes for REST error responses.
	BadRequest           // 400 Bad Request
	UnsupportedMediaType // 415 Unsupported Media Type
	NotAcceptable        // 406 Not Acceptable
	PayloadTooLarge      // 413 Payload Too Large
	TooManyRequests      // 429 Too Many Requests
	UnprocessableContent // 422 Unprocessable Content

	_maxCode
)

var strToCode = map[string]Code{
	`"OK"`:          OK,
	`"OKNoContent"`: OK,
	`"CANCELLED"`:/* [sic] */ Canceled,
	`"UNKNOWN"`:                Unknown,
	`"INVALID_ARGUMENT"`:       InvalidArgument,
	`"DEADLINE_EXCEEDED"`:      DeadlineExceeded,
	`"NOT_FOUND"`:              NotFound,
	`"ALREADY_EXISTS"`:         AlreadyExists,
	`"PERMISSION_DENIED"`:      PermissionDenied,
	`"RESOURCE_EXHAUSTED"`:     ResourceExhausted,
	`"FAILED_PRECONDITION"`:    FailedPrecondition,
	`"ABORTED"`:                Aborted,
	`"OUT_OF_RANGE"`:           OutOfRange,
	`"UNIMPLEMENTED"`:          Unimplemented,
	`"INTERNAL"`:               Internal,
	`"UNAVAILABLE"`:            Unavailable,
	`"DATA_LOSS"`:              DataLoss,
	`"UNAUTHENTICATED"`:        Unauthenticated,
	`"BAD_REQUEST"`:            BadRequest,
	`"UNSUPPORTED_MEDIA_TYPE"`: UnsupportedMediaType,
	`"NOT_ACCEPTABLE"`:         NotAcceptable,
	`"PAYLOAD_TOO_LARGE"`:      PayloadTooLarge,
	`"TOO_MANY_REQUESTS"`:      TooManyRequests,
	`"UNPROCESSABLE_CONTENT"`:  UnprocessableContent,
}

func HTTPStatusFromCode(code Code) int {
	switch code {
	case OK:
		return http.StatusOK
	case OkNoContent:
		return http.StatusNoContent
	case Canceled:
		return 499
	case Unknown:
		return http.StatusInternalServerError
	case InvalidArgument:
		return http.StatusBadRequest
	case DeadlineExceeded:
		return http.StatusGatewayTimeout
	case NotFound:
		return http.StatusNotFound
	case AlreadyExists:
		return http.StatusConflict
	case PermissionDenied:
		return http.StatusForbidden
	case Unauthenticated:
		return http.StatusUnauthorized
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case Aborted:
		return http.StatusConflict
	case OutOfRange:
		return http.StatusBadRequest
	case Unimplemented:
		return http.StatusNotImplemented
	case Internal:
		return http.StatusInternalServerError
	case Unavailable:
		return http.StatusServiceUnavailable
	case DataLoss:
		return http.StatusInternalServerError
	case BadRequest:
		return http.StatusBadRequest
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	case NotAcceptable:
		return http.StatusNotAcceptable
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case TooManyRequests:
		return http.StatusTooManyRequests
	case UnprocessableContent:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case OkNoContent:
		return "OKNoContent"
	case Canceled:
		return "Canceled"
	case Unknown:
		return "Unknown"
	case InvalidArgument:
		return "InvalidArgument"
	case DeadlineExceeded:
		return "DeadlineExceeded"
	case NotFound:
		return "NotFound"
	case AlreadyExists:
		return "AlreadyExists"
	case PermissionDenied:
		return "PermissionDenied"
	case ResourceExhausted:
		return "ResourceExhausted"
	case FailedPrecondition:
		return "FailedPrecondition"
	case Aborted:
		return "Aborted"
	case OutOfRange:
		return "OutOfRange"
	case Unimplemented:
		return "Unimplemented"
	case Internal:
		return "Internal"
	case Unavailable:
		return "Unavailable"
	case DataLoss:
		return "DataLoss"
	case Unauthenticated:
		return "Unauthenticated"
	case BadRequest:
		return "BadRequest"
	case UnsupportedMediaType:
		return "UnsupportedMediaType"
	case NotAcceptable:
		return "NotAcceptable"
	case PayloadTooLarge:
		return "PayloadTooLarge"
	case TooManyRequests:
		return "TooManyRequests"
	case UnprocessableContent:
		return "UnprocessableContent"
	default:
		return "Code(" + strconv.FormatInt(int64(c), 10) + ")"
	}
}
