package client

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// RateLimiter implements a token-bucket limiter backed by golang.org/x/time/rate
// while retaining the previous interval/next fields for limited backwards
// compatibility with existing tests and callers that inspect those fields.
type RateLimiter struct {
	// interval is kept for backwards compatibility and testing convenience.
	interval time.Duration

	// limiter is the actual token-bucket limiter used for runtime throttling.
	limiter *rate.Limiter

	mu   sync.Mutex
	next time.Time
}

func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	if requestsPerMinute <= 0 {
		return nil
	}

	interval := time.Minute / time.Duration(requestsPerMinute)

	// rate.Limiter expects events per second. Convert RPM -> RPS.
	rps := float64(requestsPerMinute) / 60.0

	// Burst set to 1 to emulate a steady rate with minimal bursting.
	limiter := rate.NewLimiter(rate.Limit(rps), 1)

	return &RateLimiter{
		interval: interval,
		limiter:  limiter,
	}
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	if r == nil {
		return nil
	}

	// If no underlying limiter is configured, fall back to no-op behavior.
	if r.limiter == nil {
		return nil
	}

	return r.limiter.Wait(ctx)
}
