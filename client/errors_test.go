package client

import "testing"

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
}

func TestIsTemporary(t *testing.T) {
	if !IsTemporary(&APIError{Code: errorCodeTooManyRequests}) {
		t.Fatal("expected APIError code 5 to be temporary")
	}

	if !IsTemporary(&HTTPError{StatusCode: 503}) {
		t.Fatal("expected HTTP 503 to be temporary")
	}

	if IsTemporary(&APIError{Code: 2}) {
		t.Fatal("did not expect invalid key errors to be temporary")
	}
}
