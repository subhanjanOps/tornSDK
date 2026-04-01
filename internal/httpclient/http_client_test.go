package httpclient

import (
	"context"
	"io"
	"net/http"
	"testing"
)

type testRoundTripFunc func(*http.Request) (*http.Response, error)

func (f testRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

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

func TestNewRequestAndCloneHelpers(t *testing.T) {
	req := NewRequest(http.MethodPost, "/user/bars")
	if got, want := req.Path, "user/bars"; got != want {
		t.Fatalf("unexpected trimmed path: got %q want %q", got, want)
	}

	var nilReq *Request
	if cloned := nilReq.Clone(); cloned != nil {
		t.Fatalf("expected nil clone, got %#v", cloned)
	}
}

func TestAddQuerySetQueryAndSelectionsHelpers(t *testing.T) {
	var nilReq *Request
	if got := nilReq.AddQuery("key", "value"); got != nil {
		t.Fatalf("expected nil AddQuery receiver to stay nil, got %#v", got)
	}

	if got := nilReq.SetQuery("key", "value"); got != nil {
		t.Fatalf("expected nil SetQuery receiver to stay nil, got %#v", got)
	}

	req := &Request{}
	req.SetQuery("init", "value")
	if got, want := req.Query.Get("init"), "value"; got != want {
		t.Fatalf("unexpected initialized query value: got %q want %q", got, want)
	}

	req = &Request{}
	req.AddQuery("init", "value")
	if got, want := req.Query.Get("init"), "value"; got != want {
		t.Fatalf("unexpected initialized add query value: got %q want %q", got, want)
	}

	req.AddQuery("a", "1", "2")
	if got, want := req.Query["a"][0], "1"; got != want {
		t.Fatalf("unexpected first added query value: got %q want %q", got, want)
	}

	req.SetQuery("a", "3")
	if got, want := req.Query.Get("a"), "3"; got != want {
		t.Fatalf("unexpected set query value: got %q want %q", got, want)
	}

	req.SetSelections(" bars ", "", "profile")
	if got, want := req.Query.Get("selections"), "bars,profile"; got != want {
		t.Fatalf("unexpected selections query: got %q want %q", got, want)
	}

	req.SetSelections(" ", "")
	if got := req.Query.Get("selections"); got != "" {
		t.Fatalf("expected selections query to be removed, got %q", got)
	}

	req.SetQuery("a")
	if got := req.Query.Get("a"); got != "" {
		t.Fatalf("expected query key to be removed, got %q", got)
	}
}

func TestBuildRequestNilDefaultAndInvalid(t *testing.T) {
	transport := New("https://api.torn.com/v2/", &http.Client{}, "")

	if _, err := transport.BuildRequest(context.Background(), nil); err == nil {
		t.Fatal("expected nil request error")
	}

	req := &Request{
		Path: "user/bars",
	}

	httpReq, err := transport.BuildRequest(context.Background(), req)
	if err != nil {
		t.Fatalf("BuildRequest returned error: %v", err)
	}

	if got, want := httpReq.Method, http.MethodGet; got != want {
		t.Fatalf("unexpected default method: got %q want %q", got, want)
	}

	if got := httpReq.Header.Get("User-Agent"); got != "" {
		t.Fatalf("expected empty user agent when not configured, got %q", got)
	}

	badReq := NewRequest(http.MethodGet, "bad\npath")
	if _, err := transport.BuildRequest(context.Background(), badReq); err == nil {
		t.Fatal("expected invalid URL error")
	}
}

func TestClientDoUsesUnderlyingHTTPClient(t *testing.T) {
	called := false
	client := New("https://api.torn.com/v2", &http.Client{
		Transport: testRoundTripFunc(func(req *http.Request) (*http.Response, error) {
			called = true
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Header:     make(http.Header),
				Body:       io.NopCloser(req.Body),
			}, nil
		}),
	}, "agent")

	req, err := client.BuildRequest(context.Background(), NewRequest(http.MethodGet, "user/bars"))
	if err != nil {
		t.Fatalf("BuildRequest returned error: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}

	if !called {
		t.Fatal("expected underlying HTTP client to be used")
	}

	if resp == nil {
		t.Fatal("expected response")
	}
}
