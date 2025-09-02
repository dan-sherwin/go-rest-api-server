package apiresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiResponse represents a standard structure for API responses, containing fields for response code, message, description, and optional details. Fields isNil and httpCode are used internally for response handling and are omitted from the JSON output.
type (
	ApiResponse struct {
		Code        string `json:"code"`
		Message     string `json:"message"`
		Description string `json:"description"`
		Details     any    `json:"details"`
		isNil       bool   `json:"-"`
		httpCode    int    `json:"-"`
	}
)

// EAPI represents an enumerated type for defining custom API error codes. It is utilized to standardize error handling and map error types to HTTP response codes in an application.
//
//go:generate stringer -type=EAPI
type EAPI int

// EAPIError represents a generic API error.
// EAPIMFARequired indicates that Multi-Factor Authentication is required.
// EAPIInvalidCredentials signifies invalid user credentials.
// EAPIPathNotFound denotes that the requested API path is not found.
// EAPIFieldValidation indicates field validation errors.
// EAPIPasswordResetFail represents failure in password reset process.
// EAPIJSONBad signifies an invalid or malformed JSON error.
// EAPIDenied represents a denied access error.
// EIAMError denotes an Identity and Access Management system error.
const (
	EAPIError EAPI = iota
	EAPIMFARequired
	EAPIInvalidCredentials
	EAPIPathNotFound
	EAPIFieldValidation
	EAPIPasswordResetFail
	EAPIJSONBad
	EAPIDenied
	EIAMError
)

// responseCodeMap maps EAPI error codes to corresponding HTTP status codes to standardize API responses.
var responseCodeMap = map[EAPI]int{
	EAPIError:              http.StatusInternalServerError,
	EAPIMFARequired:        http.StatusUnauthorized,
	EAPIInvalidCredentials: http.StatusUnauthorized,
	EAPIPathNotFound:       http.StatusNotFound,
	EAPIFieldValidation:    http.StatusBadRequest,
	EAPIPasswordResetFail:  http.StatusInternalServerError,
	EAPIJSONBad:            http.StatusBadRequest,
	EAPIDenied:             http.StatusForbidden,
	EIAMError:              http.StatusBadRequest,
}

// IsNil checks if the ApiResponse instance is marked as nil based on its internal state.
func (ar *ApiResponse) IsNil() bool {
	return ar.isNil
}

// Send sends an HTTP response using the provided gin.Context based on the ApiResponse.
// It returns "success" for nil responses with StatusOK, sends ar.Message for zero httpCode with InternalServerError,
// otherwise sends the ApiResponse with the specified httpCode.
func (ar *ApiResponse) Send(c *gin.Context) {
	if ar.IsNil() {
		c.JSON(http.StatusOK, "success")
		return
	}
	if ar.httpCode == 0 {
		c.JSON(http.StatusInternalServerError, ar.Message)
		return
	}
	c.JSON(ar.httpCode, ar)
}

// Nil creates and returns an ApiResponse with the isNil field set to true, representing an empty or uninitialized response.
func Nil() ApiResponse {
	return ApiResponse{
		isNil: true,
	}
}

// New creates a new ApiResponse instance with the given EAPI code, message, and optional details, mapping the EAPI code to its respective HTTP status code.
func New(respCode EAPI, message string, details ...any) ApiResponse {
	resp := ApiResponse{
		Code:     respCode.String(),
		Message:  message,
		Details:  details,
		httpCode: responseCodeMap[respCode],
	}
	if details != nil {
		resp.Details = details[0]
		if len(details) > 1 {
			resp.Details = details
		}
	}
	return resp
}

// SendApiErrorResponse sends a JSON error response with a status code determined by the EAPI value, including an error message and optional details.
func SendApiErrorResponse(c *gin.Context, respCode EAPI, message string, details ...any) {
	c.JSON(responseCodeMap[respCode], New(respCode, message, details...))
}

// SendApiResponse sends a JSON response with HTTP status 200 (OK) using the provided context and data.
func SendApiResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
