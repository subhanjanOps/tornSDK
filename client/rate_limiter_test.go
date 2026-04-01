package client

import (
	"context"
	"errors"
	"testing"
	"time"
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
		t.Fatalf("expected zero interval limiter to no-op, got %v", err)
	}
}

func TestRateLimiterWaitImmediateAndPastSchedule(t *testing.T) {
	limiter := &RateLimiter{interval: time.Millisecond}
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected first wait to be immediate, got %v", err)
	}

	limiter.next = time.Now().Add(-time.Second)
	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected past schedule to be immediate, got %v", err)
	}
}

func TestRateLimiterWaitDelayAndCancel(t *testing.T) {
	limiter := &RateLimiter{
		interval: 5 * time.Millisecond,
		next:     time.Now().Add(2 * time.Millisecond),
	}

	if err := limiter.Wait(context.Background()); err != nil {
		t.Fatalf("expected delayed wait to succeed, got %v", err)
	}

	limiter.next = time.Now().Add(time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := limiter.Wait(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation error, got %v", err)
	}
}
