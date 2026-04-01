package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type stubLimiter struct {
	calls int
	err   error
}

func (l *stubLimiter) Wait(context.Context) error {
	l.calls++
	return l.err
}

type errReadCloser struct{}

func (errReadCloser) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func (errReadCloser) Close() error {
	return nil
}

func newHTTPResponse(statusCode int, status, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Status:     status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestClientDoNilClientAndRequest(t *testing.T) {
	var nilClient *Client
	err := nilClient.Do(context.Background(), NewRequest(http.MethodGet, "user/bars"), nil)
	if err == nil || err.Error() != "nil client" {
		t.Fatalf("expected nil client error, got %v", err)
	}

	client := &Client{}
	err = client.Do(context.Background(), nil, nil)
	if err == nil || err.Error() != "nil request" {
		t.Fatalf("expected nil request error, got %v", err)
	}
}

func TestClientDoReturnsLimiterError(t *testing.T) {
	limiter := &stubLimiter{err: errors.New("rate limited locally")}
	client := &Client{
		http:        httpclient.New("https://example.com", &http.Client{}, ""),
		limiter:     limiter,
		retryPolicy: RetryPolicy{},
	}

	err := client.Do(context.Background(), NewRequest(http.MethodGet, "user/bars"), nil)
	if err == nil || err.Error() != "rate limited locally" {
		t.Fatalf("expected limiter error, got %v", err)
	}

	if got, want := limiter.calls, 1; got != want {
		t.Fatalf("unexpected limiter calls: got %d want %d", got, want)
	}
}

func TestClientDoReturnsBuildRequestError(t *testing.T) {
	client := &Client{
		http:        httpclient.New("https://example.com", &http.Client{}, ""),
		retryPolicy: RetryPolicy{},
	}

	err := client.Do(context.Background(), NewRequest(http.MethodGet, "bad\npath"), nil)
	if err == nil {
		t.Fatal("expected build request error")
	}
}

func TestClientDoRetriesTemporaryHTTPErrorAndSetsKey(t *testing.T) {
	attempts := 0
	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			attempts++

			if got, want := req.URL.Query().Get("key"), "secret"; got != want {
				t.Fatalf("unexpected API key query: got %q want %q", got, want)
			}

			if attempts == 1 {
				resp := newHTTPResponse(http.StatusServiceUnavailable, "503 Service Unavailable", "busy")
				resp.Header.Set("Retry-After", "0")
				return resp, nil
			}

			return newHTTPResponse(http.StatusOK, "200 OK", `{"ok":true}`), nil
		}),
	}

	client := &Client{
		apiKey:      "secret",
		http:        httpclient.New("https://example.com", httpClient, ""),
		retryPolicy: NewRetryPolicy(1, time.Millisecond, time.Millisecond),
	}

	var out struct {
		OK bool `json:"ok"`
	}

	if err := client.Do(context.Background(), NewRequest(http.MethodGet, "user/bars"), &out); err != nil {
		t.Fatalf("Do returned error: %v", err)
	}

	if !out.OK {
		t.Fatal("expected decoded success response")
	}

	if got, want := attempts, 2; got != want {
		t.Fatalf("unexpected attempts: got %d want %d", got, want)
	}
}

func TestClientDoReturnsNonTemporaryHTTPErrorWithoutRetry(t *testing.T) {
	attempts := 0
	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			attempts++
			return newHTTPResponse(http.StatusBadRequest, "400 Bad Request", "bad request"), nil
		}),
	}

	client := &Client{
		http:        httpclient.New("https://example.com", httpClient, ""),
		retryPolicy: NewRetryPolicy(3, time.Millisecond, time.Millisecond),
	}

	err := client.Do(context.Background(), NewRequest(http.MethodGet, "user/bars"), nil)
	if err == nil {
		t.Fatal("expected HTTP error")
	}

	if got, want := attempts, 1; got != want {
		t.Fatalf("unexpected attempts: got %d want %d", got, want)
	}
}

func TestClientDoReturnsContextErrorDuringRetrySleep(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	attempts := 0
	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			attempts++
			cancel()
			resp := newHTTPResponse(http.StatusServiceUnavailable, "503 Service Unavailable", "busy")
			resp.Header.Set("Retry-After", "1")
			return resp, nil
		}),
	}

	client := &Client{
		http:        httpclient.New("https://example.com", httpClient, ""),
		retryPolicy: NewRetryPolicy(1, time.Millisecond, time.Millisecond),
	}

	err := client.Do(ctx, NewRequest(http.MethodGet, "user/bars"), nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation during retry sleep, got %v", err)
	}

	if got, want := attempts, 1; got != want {
		t.Fatalf("unexpected attempts: got %d want %d", got, want)
	}
}

func TestDecodeResponse(t *testing.T) {
	t.Run("nil response", func(t *testing.T) {
		if err := decodeResponse(nil, nil); err == nil || err.Error() != "nil response" {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("read error", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Body:       errReadCloser{},
		}

		err := decodeResponse(resp, nil)
		if err == nil || !strings.Contains(err.Error(), "read response body") {
			t.Fatalf("expected read error, got %v", err)
		}
	})

	t.Run("api error", func(t *testing.T) {
		resp := newHTTPResponse(http.StatusOK, "200 OK", `{"error":{"code":5,"error":"Too many requests"}}`)
		err := decodeResponse(resp, nil)

		apiErr, ok := err.(*APIError)
		if !ok {
			t.Fatalf("expected APIError, got %T", err)
		}

		if got, want := apiErr.Code, 5; got != want {
			t.Fatalf("unexpected API error code: got %d want %d", got, want)
		}
	})

	t.Run("http error", func(t *testing.T) {
		resp := newHTTPResponse(http.StatusBadGateway, "502 Bad Gateway", "backend failed")
		err := decodeResponse(resp, nil)

		httpErr, ok := err.(*HTTPError)
		if !ok {
			t.Fatalf("expected HTTPError, got %T", err)
		}

		if got, want := httpErr.StatusCode, http.StatusBadGateway; got != want {
			t.Fatalf("unexpected status code: got %d want %d", got, want)
		}
	})

	t.Run("nil value or blank body", func(t *testing.T) {
		if err := decodeResponse(newHTTPResponse(http.StatusOK, "200 OK", ""), nil); err != nil {
			t.Fatalf("expected nil error for blank body, got %v", err)
		}
	})

	t.Run("decode error", func(t *testing.T) {
		resp := newHTTPResponse(http.StatusOK, "200 OK", "{")
		err := decodeResponse(resp, &struct{}{})
		if err == nil || !strings.Contains(err.Error(), "decode response body") {
			t.Fatalf("expected decode error, got %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		resp := newHTTPResponse(http.StatusOK, "200 OK", `{"value":7}`)
		var out struct {
			Value int `json:"value"`
		}

		if err := decodeResponse(resp, &out); err != nil {
			t.Fatalf("expected success, got %v", err)
		}

		if got, want := out.Value, 7; got != want {
			t.Fatalf("unexpected decoded value: got %d want %d", got, want)
		}
	})
}
