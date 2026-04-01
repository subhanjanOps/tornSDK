package client

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestNewRetryPolicyDefaultsAndClamp(t *testing.T) {
	policy := NewRetryPolicy(-1, 0, 0)
	if got, want := policy.MaxRetries, 0; got != want {
		t.Fatalf("unexpected max retries: got %d want %d", got, want)
	}

	if got, want := policy.MinWait, defaultRetryWaitMin; got != want {
		t.Fatalf("unexpected min wait: got %s want %s", got, want)
	}

	if got, want := policy.MaxWait, defaultRetryWaitMax; got != want {
		t.Fatalf("unexpected max wait: got %s want %s", got, want)
	}

	policy = NewRetryPolicy(1, 3*time.Second, time.Second)
	if got, want := policy.MaxWait, 3*time.Second; got != want {
		t.Fatalf("unexpected clamped max wait: got %s want %s", got, want)
	}
}

func TestRetryPolicyNextBackoffStopsWhenDone(t *testing.T) {
	policy := NewRetryPolicy(1, time.Second, 2*time.Second)

	if wait, ok := policy.NextBackoff(1, nil, &HTTPError{StatusCode: http.StatusServiceUnavailable}); ok || wait != 0 {
		t.Fatalf("expected retries to stop after max retries, got wait=%s ok=%v", wait, ok)
	}

	if wait, ok := policy.NextBackoff(0, nil, errors.New("permanent")); ok || wait != 0 {
		t.Fatalf("expected permanent errors to stop retries, got wait=%s ok=%v", wait, ok)
	}
}

func TestRetryPolicyNextBackoffUsesRetryAfterAndBackoff(t *testing.T) {
	policy := NewRetryPolicy(3, time.Second, 4*time.Second)
	resp := &http.Response{
		Header: http.Header{"Retry-After": []string{"2"}},
	}

	wait, ok := policy.NextBackoff(0, resp, &HTTPError{StatusCode: http.StatusServiceUnavailable})
	if !ok || wait != 2*time.Second {
		t.Fatalf("expected Retry-After seconds to be used, got wait=%s ok=%v", wait, ok)
	}

	future := time.Now().Add(1500 * time.Millisecond).UTC().Format(http.TimeFormat)
	resp.Header.Set("Retry-After", future)

	wait, ok = policy.NextBackoff(0, resp, &HTTPError{StatusCode: http.StatusServiceUnavailable})
	if !ok || wait < 0 || wait > 2*time.Second {
		t.Fatalf("expected Retry-After date to be used, got wait=%s ok=%v", wait, ok)
	}

	resp.Header.Set("Retry-After", "bad-value")
	wait, ok = policy.NextBackoff(2, resp, &HTTPError{StatusCode: http.StatusServiceUnavailable})
	if !ok || wait != 4*time.Second {
		t.Fatalf("expected exponential backoff fallback, got wait=%s ok=%v", wait, ok)
	}
}

func TestRetryPolicyBackoffCapsAtMax(t *testing.T) {
	policy := NewRetryPolicy(5, time.Second, 3*time.Second)

	if got, want := policy.backoff(0), time.Second; got != want {
		t.Fatalf("unexpected backoff for attempt 0: got %s want %s", got, want)
	}

	if got, want := policy.backoff(1), 2*time.Second; got != want {
		t.Fatalf("unexpected backoff for attempt 1: got %s want %s", got, want)
	}

	if got, want := policy.backoff(3), 3*time.Second; got != want {
		t.Fatalf("unexpected backoff cap: got %s want %s", got, want)
	}

	invalidPolicy := RetryPolicy{MinWait: 5 * time.Second, MaxWait: 3 * time.Second}
	if got, want := invalidPolicy.backoff(0), 3*time.Second; got != want {
		t.Fatalf("unexpected backoff for invalid policy: got %s want %s", got, want)
	}
}

func TestParseRetryAfter(t *testing.T) {
	if wait, ok := parseRetryAfter(""); ok || wait != 0 {
		t.Fatalf("expected empty Retry-After to fail, got wait=%s ok=%v", wait, ok)
	}

	if wait, ok := parseRetryAfter("-3"); !ok || wait != 0 {
		t.Fatalf("expected negative Retry-After seconds to clamp to zero, got wait=%s ok=%v", wait, ok)
	}

	past := time.Now().Add(-time.Second).UTC().Format(http.TimeFormat)
	if wait, ok := parseRetryAfter(past); !ok || wait != 0 {
		t.Fatalf("expected past Retry-After date to clamp to zero, got wait=%s ok=%v", wait, ok)
	}

	if wait, ok := parseRetryAfter("not-a-date"); ok || wait != 0 {
		t.Fatalf("expected invalid Retry-After to fail, got wait=%s ok=%v", wait, ok)
	}
}

func TestSleepContext(t *testing.T) {
	if err := sleepContext(context.Background(), 0); err != nil {
		t.Fatalf("expected zero wait to succeed, got %v", err)
	}

	if err := sleepContext(context.Background(), time.Millisecond); err != nil {
		t.Fatalf("expected timer wait to succeed, got %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := sleepContext(ctx, time.Second); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
}
