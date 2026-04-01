package client

import (
	"errors"
	"strings"
	"testing"
)

type temporaryNetError struct {
	timeout   bool
	temporary bool
}

func (e temporaryNetError) Error() string   { return "temporary net error" }
func (e temporaryNetError) Timeout() bool   { return e.timeout }
func (e temporaryNetError) Temporary() bool { return e.temporary }

func TestParseAPIError(t *testing.T) {
	body := []byte(`{"error":{"code":5,"error":"Too many requests"}}`)

	apiErr := parseAPIError(body)
	if apiErr == nil {
		t.Fatal("expected api error")
	}

	if apiErr.Code != 5 {
		t.Fatalf("unexpected error code: got %d", apiErr.Code)
	}

	if apiErr.Message != "Too many requests" {
		t.Fatalf("unexpected error message: got %q", apiErr.Message)
	}

	if parseAPIError([]byte("{")) != nil {
		t.Fatal("expected invalid JSON to return nil")
	}

	if parseAPIError([]byte(`{"error":null}`)) != nil {
		t.Fatal("expected nil error envelope to return nil")
	}

	if parseAPIError([]byte(`{"error":{"code":0,"error":""}}`)) != nil {
		t.Fatal("expected empty API error to return nil")
	}
}

func TestIsTemporary(t *testing.T) {
	if IsTemporary(nil) {
		t.Fatal("did not expect nil to be temporary")
	}

	if !IsTemporary(&APIError{Code: errorCodeTooManyRequests}) {
		t.Fatal("expected APIError code 5 to be temporary")
	}

	if !IsTemporary(&HTTPError{StatusCode: 503}) {
		t.Fatal("expected HTTP 503 to be temporary")
	}

	if IsTemporary(&APIError{Code: 2}) {
		t.Fatal("did not expect invalid key errors to be temporary")
	}

	if !IsTemporary(temporaryNetError{timeout: true}) {
		t.Fatal("expected timeout net errors to be temporary")
	}

	if IsTemporary(errors.New("plain error")) {
		t.Fatal("did not expect plain errors to be temporary")
	}
}

func TestAPIErrorErrorAndTemporary(t *testing.T) {
	var nilErr *APIError
	if got := nilErr.Error(); got != "" {
		t.Fatalf("expected nil APIError string to be empty, got %q", got)
	}

	err := &APIError{Code: 2}
	if got, want := err.Error(), "torn api error code 2"; got != want {
		t.Fatalf("unexpected APIError string: got %q want %q", got, want)
	}

	err.Message = "Invalid key"
	if got, want := err.Error(), "torn api error 2: Invalid key"; got != want {
		t.Fatalf("unexpected APIError string: got %q want %q", got, want)
	}

	if (*APIError)(nil).Temporary() {
		t.Fatal("did not expect nil APIError to be temporary")
	}
}

func TestHTTPErrorErrorAndTemporary(t *testing.T) {
	var nilErr *HTTPError
	if got := nilErr.Error(); got != "" {
		t.Fatalf("expected nil HTTPError string to be empty, got %q", got)
	}

	err := &HTTPError{StatusCode: 400, Status: "400 Bad Request"}
	if got, want := err.Error(), "unexpected http status: 400 Bad Request"; got != want {
		t.Fatalf("unexpected HTTPError string: got %q want %q", got, want)
	}

	err.Body = "bad request"
	if got, want := err.Error(), "unexpected http status: 400 Bad Request: bad request"; got != want {
		t.Fatalf("unexpected HTTPError string: got %q want %q", got, want)
	}

	if (*HTTPError)(nil).Temporary() {
		t.Fatal("did not expect nil HTTPError to be temporary")
	}

	if !(&HTTPError{StatusCode: 429}).Temporary() {
		t.Fatal("expected 429 to be temporary")
	}
}

func TestShortenBody(t *testing.T) {
	short := []byte(" short ")
	if got, want := shortenBody(short), "short"; got != want {
		t.Fatalf("unexpected shortened short body: got %q want %q", got, want)
	}

	long := []byte(strings.Repeat("a", 300))
	got := shortenBody(long)
	if !strings.HasSuffix(got, "...") {
		t.Fatalf("expected shortened long body to end with ellipsis, got %q", got)
	}

	if gotLen, wantLen := len(got), 259; gotLen != wantLen {
		t.Fatalf("unexpected shortened long body length: got %d want %d", gotLen, wantLen)
	}
}
