package client

import (
	"testing"
)

func TestParseAPIError_PlainString(t *testing.T) {
	body := []byte(`{"error":"simple message"}`)
	apiErr := parseAPIError(body)
	if apiErr == nil {
		t.Fatal("expected APIError")
	}
	if got, want := apiErr.Message, "simple message"; got != want {
		t.Fatalf("unexpected message: got %q want %q", got, want)
	}
}

func TestParseAPIError_ErrorsArray(t *testing.T) {
	body := []byte(`{"errors":[{"code":5,"error":"Too many requests"}]}`)
	apiErr := parseAPIError(body)
	if apiErr == nil {
		t.Fatal("expected APIError")
	}
	if got, want := apiErr.Code, 5; got != want {
		t.Fatalf("unexpected code: got %d want %d", got, want)
	}
}
