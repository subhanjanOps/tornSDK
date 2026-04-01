package httpclient

import (
	"context"
	"net/http"
	"testing"
)

func TestBuildRequest(t *testing.T) {
	transport := New("https://api.torn.com/v2/", &http.Client{}, "tornSDK/test")

	req := NewRequest(http.MethodGet, "user/bars")
	req.SetQuery("key", "abc123")
	req.Headers.Set("X-Test", "1")

	httpReq, err := transport.BuildRequest(context.Background(), req)
	if err != nil {
		t.Fatalf("BuildRequest returned error: %v", err)
	}

	if got, want := httpReq.URL.String(), "https://api.torn.com/v2/user/bars?key=abc123"; got != want {
		t.Fatalf("unexpected URL: got %q want %q", got, want)
	}

	if got := httpReq.Header.Get("Accept"); got != "application/json" {
		t.Fatalf("unexpected Accept header: got %q", got)
	}

	if got := httpReq.Header.Get("User-Agent"); got != "tornSDK/test" {
		t.Fatalf("unexpected User-Agent header: got %q", got)
	}

	if got := httpReq.Header.Get("X-Test"); got != "1" {
		t.Fatalf("unexpected custom header: got %q", got)
	}
}

func TestRequestClone(t *testing.T) {
	req := NewRequest(http.MethodGet, "user/bars")
	req.SetQuery("key", "original")
	req.Headers.Set("X-Test", "1")

	cloned := req.Clone()
	cloned.SetQuery("key", "cloned")
	cloned.Headers.Set("X-Test", "2")

	if got := req.Query.Get("key"); got != "original" {
		t.Fatalf("expected original query to remain unchanged, got %q", got)
	}

	if got := req.Headers.Get("X-Test"); got != "1" {
		t.Fatalf("expected original header to remain unchanged, got %q", got)
	}
}
