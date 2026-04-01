package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

const (
	errorCodeTooManyRequests = 5
	errorCodeIPBlock         = 8
	errorCodeAPIDisabled     = 9
	errorCodeTemporary       = 15
	errorCodeBackendError    = 17
	errorCodeClosedTemporary = 24
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}

	if e.Message == "" {
		return fmt.Sprintf("torn api error code %d", e.Code)
	}

	return fmt.Sprintf("torn api error %d: %s", e.Code, e.Message)
}

func (e *APIError) Temporary() bool {
	if e == nil {
		return false
	}

	switch e.Code {
	case errorCodeTooManyRequests, errorCodeIPBlock, errorCodeAPIDisabled, errorCodeTemporary, errorCodeBackendError, errorCodeClosedTemporary:
		return true
	default:
		return false
	}
}

type HTTPError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HTTPError) Error() string {
	if e == nil {
		return ""
	}

	if e.Body == "" {
		return fmt.Sprintf("unexpected http status: %s", e.Status)
	}

	return fmt.Sprintf("unexpected http status: %s: %s", e.Status, e.Body)
}

func (e *HTTPError) Temporary() bool {
	if e == nil {
		return false
	}

	return e.StatusCode == http.StatusTooManyRequests || e.StatusCode >= http.StatusInternalServerError
}

func IsTemporary(err error) bool {
	if err == nil {
		return false
	}

	var netErr net.Error
	if errors.As(err, &netErr) && (netErr.Timeout() || netErr.Temporary()) {
		return true
	}

	var temporary interface{ Temporary() bool }
	return errors.As(err, &temporary) && temporary.Temporary()
}

func parseAPIError(body []byte) *APIError {
	var envelope struct {
		Error *APIError `json:"error"`
	}

	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil
	}

	if envelope.Error == nil {
		return nil
	}

	if envelope.Error.Code == 0 && envelope.Error.Message == "" {
		return nil
	}

	return envelope.Error
}

func shortenBody(body []byte) string {
	text := strings.TrimSpace(string(body))
	if len(text) <= 256 {
		return text
	}

	return text[:256] + "..."
}
