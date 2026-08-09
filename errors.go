package bursapay

import (
	"fmt"
)

// APIError represents an error response returned by the BursaPay API.
type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	ErrorMsg   string `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = e.ErrorMsg
	}
	if msg == "" {
		msg = "unknown error"
	}
	if e.Code != "" {
		return fmt.Sprintf("bursapay: status %d [%s]: %s", e.StatusCode, e.Code, msg)
	}
	return fmt.Sprintf("bursapay: status %d: %s", e.StatusCode, msg)
}

// AuthenticationError indicates an invalid API key or authorization failure (401 / 403).
type AuthenticationError struct {
	APIError
}

// InvalidRequestError indicates invalid parameters or validation failure (400 / 422).
type InvalidRequestError struct {
	APIError
}

// RateLimitError indicates rate limit limits reached (429).
type RateLimitError struct {
	APIError
}

// SignatureVerificationError indicates webhook signature validation failure.
type SignatureVerificationError struct {
	Message string
}

func (e *SignatureVerificationError) Error() string {
	return fmt.Sprintf("bursapay signature verification error: %s", e.Message)
}
