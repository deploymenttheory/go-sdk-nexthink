package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

// HTTP Status Codes
const (
	// Success status codes
	StatusOK      = 200 // Successful GET requests
	StatusCreated = 201 // Successful POST requests

	// Client error status codes
	StatusBadRequest          = 400 // Bad request, invalid arguments
	StatusUnauthorized        = 401 // Authentication required, invalid credentials
	StatusForbidden           = 403 // Forbidden operation
	StatusNotFound            = 404 // Resource not found
	StatusConflict            = 409 // Resource already exists
	StatusUnprocessableEntity = 422 // Validation errors
	StatusTooManyRequests     = 429 // Rate limit exceeded

	// Server error status codes
	StatusInternalServerError = 500 // Server-side error
	StatusBadGateway          = 502 // Gateway error
	StatusServiceUnavailable  = 503 // Service temporarily unavailable
	StatusGatewayTimeout      = 504 // Deadline exceeded
)

// APIError represents an error response from the Nexthink API
type APIError struct {
	Code    string `json:"code,omitempty"`    // Error code if provided
	Message string `json:"message,omitempty"` // Error message
	Details string `json:"details,omitempty"` // Additional error details

	// HTTP response details
	StatusCode int    // HTTP status code
	Status     string // HTTP status text
	Endpoint   string // API endpoint that returned the error
	Method     string // HTTP method used
}

// genericErrorResponse represents a generic API error response wrapper
type genericErrorResponse struct {
	Error   *APIError       `json:"error,omitempty"`
	Message string          `json:"message,omitempty"`
	Details json.RawMessage `json:"details,omitempty"`
	Code    json.RawMessage `json:"code,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("Nexthink API error (%d %s) [%s] at %s %s: %s",
			e.StatusCode, e.Status, e.Code, e.Method, e.Endpoint, e.Message)
	}
	return fmt.Sprintf("Nexthink API error (%d %s) at %s %s: %s",
		e.StatusCode, e.Status, e.Method, e.Endpoint, e.Message)
}

// ParseErrorResponse parses an error response from the API
func ParseErrorResponse(body []byte, statusCode int, status, method, endpoint string, logger *zap.Logger) error {
	apiError := &APIError{
		StatusCode: statusCode,
		Status:     status,
		Endpoint:   endpoint,
		Method:     method,
	}

	// Try to parse as structured error response
	var errResp genericErrorResponse
	if err := json.Unmarshal(body, &errResp); err == nil {
		// Check if error object exists
		if errResp.Error != nil {
			apiError.Code = errResp.Error.Code
			apiError.Message = errResp.Error.Message
			apiError.Details = errResp.Error.Details
		} else {
			// Try top-level message and code
			apiError.Details = decodeErrorValue(errResp.Details)
			if errResp.Message == "" {
				apiError.Message = apiError.Details
			}
			if errResp.Message != "" {
				apiError.Message = errResp.Message
			}
			if len(errResp.Code) != 0 {
				apiError.Code = decodeErrorValue(errResp.Code)
			}
		}

		if apiError.Code != "" || apiError.Message != "" {
			logger.Error("API error response",
				zap.Int("status_code", statusCode),
				zap.String("status", status),
				zap.String("method", method),
				zap.String("endpoint", endpoint),
				zap.String("error_code", apiError.Code))
			return apiError
		}
	}

	// If JSON parsing fails or doesn't match expected format, use raw body as message
	apiError.Message = string(body)
	if apiError.Message == "" {
		apiError.Message = getDefaultErrorMessage(statusCode)
	}

	logger.Error("API error response",
		zap.Int("status_code", statusCode),
		zap.String("status", status),
		zap.String("method", method),
		zap.String("endpoint", endpoint))

	return apiError
}

// getDefaultErrorMessage returns a default error message based on status code
func getDefaultErrorMessage(statusCode int) string {
	switch statusCode {
	case StatusBadRequest:
		return "The API request is invalid or malformed"
	case StatusUnauthorized:
		return "Authentication required or invalid credentials"
	case StatusForbidden:
		return "You are not allowed to perform the requested operation"
	case StatusNotFound:
		return "The requested resource was not found"
	case StatusConflict:
		return "The resource already exists"
	case StatusUnprocessableEntity:
		return "Validation error"
	case StatusTooManyRequests:
		return "Rate limit exceeded. Too many requests in a given time period"
	case StatusInternalServerError:
		return "Internal server error"
	case StatusBadGateway:
		return "Bad gateway"
	case StatusServiceUnavailable:
		return "Service temporarily unavailable. Retry might work"
	case StatusGatewayTimeout:
		return "The operation took too long to complete"
	default:
		return "Unknown error"
	}
}

// Error type check helpers

// IsBadRequest checks if the error is a bad request error (400)
func IsBadRequest(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusBadRequest
	}
	return false
}

// IsUnauthorized checks if the error is an authentication error (401)
func IsUnauthorized(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusUnauthorized
	}
	return false
}

// IsForbidden checks if the error is a forbidden error (403)
func IsForbidden(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusForbidden
	}
	return false
}

// IsNotFound checks if the error is a not found error (404)
func IsNotFound(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusNotFound
	}
	return false
}

// IsConflict checks if the error is a conflict error (409) - resource already exists
func IsConflict(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusConflict
	}
	return false
}

// IsValidationError checks if the error is a validation/unprocessable entity error (422)
func IsValidationError(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusUnprocessableEntity
	}
	return false
}

// IsRateLimited checks if the error is a rate limit error (429)
func IsRateLimited(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusTooManyRequests
	}
	return false
}

// IsServerError checks if the error is a server error (5xx)
func IsServerError(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode >= 500 && apiErr.StatusCode < 600
	}
	return false
}

// IsTransient checks if the error is transient and can be retried
func IsTransient(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode == StatusServiceUnavailable ||
			apiErr.StatusCode == StatusGatewayTimeout
	}
	return false
}

// GetErrorCode returns the error code from the error
func GetErrorCode(err error) string {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code
	}
	return ""
}

func decodeErrorValue(raw json.RawMessage) string {
	if len(raw) == 0 || string(raw) == "null" {
		return ""
	}
	var value string
	if json.Unmarshal(raw, &value) == nil {
		return value
	}
	return string(raw)
}

// UnmarshalJSON accepts numeric NQL codes as well as string service codes.
func (e *APIError) UnmarshalJSON(data []byte) error {
	var wire struct {
		Code    json.RawMessage `json:"code"`
		Message string          `json:"message"`
		Details json.RawMessage `json:"details"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	e.Code = decodeErrorValue(wire.Code)
	e.Message = wire.Message
	e.Details = decodeErrorValue(wire.Details)
	return nil
}
