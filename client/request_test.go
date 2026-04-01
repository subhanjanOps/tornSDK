package client

import (
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest(http.MethodGet, "/user/bars")
	if req == nil {
		t.Fatal("expected request")
	}

	if got, want := req.Method, http.MethodGet; got != want {
		t.Fatalf("unexpected method: got %q want %q", got, want)
	}

	if got, want := req.Path, "user/bars"; got != want {
		t.Fatalf("unexpected path: got %q want %q", got, want)
	}
}
