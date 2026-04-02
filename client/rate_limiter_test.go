package client

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestNewRateLimiter(t *testing.T) {
	if limiter := NewRateLimiter(0); limiter != nil {
		t.Fatal("expected nil limiter for zero RPM")
	}

	limiter := NewRateLimiter(60)
	if limiter == nil {
		t.Fatal("expected limiter")
	}

	if got, want := limiter.interval, time.Second; got != want {
		t.Fatalf("unexpected interval: got %s want %s", got, want)
	}
}

func TestRateLimiterWaitNoopCases(t *testing.T) {
	var nilLimiter *RateLimiter
	if err := nilLimiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected nil limiter to no-op, got %v", err)
	}

	limiter := &RateLimiter{}
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected zero-limiter to no-op, got %v", err)
	}
}

func TestRateLimiterWaitImmediateAndSubsequent(t *testing.T) {
	limiter := NewRateLimiter(600) // high RPS so waits are near-immediate
	if limiter == nil {
		t.Fatal("expected limiter")
	}

	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected first wait to be immediate, got %v", err)
	}

	// subsequent immediate call should also succeed (burst=1 may allow immediate)
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected second wait to be immediate, got %v", err)
	}
}

func TestRateLimiterWaitDelayAndCancel(t *testing.T) {
	// Create a limiter that will block on Wait by using rate 0 and burst 0.
	limiter := &RateLimiter{
		interval: 5 * time.Millisecond,
		limiter:  rate.NewLimiter(0, 1),
	}

	// consume the initial token so that Wait will block
	limiter.limiter.Allow()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := limiter.Wait(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation error, got %v", err)
	}
}
