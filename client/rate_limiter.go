package client

import (
	"context"
	"sync"
	"time"
)

type RateLimiter struct {
	interval time.Duration

	mu   sync.Mutex
	next time.Time
}

func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	if requestsPerMinute <= 0 {
		return nil
	}

	return &RateLimiter{
		interval: time.Minute / time.Duration(requestsPerMinute),
	}
}

func (r *RateLimiter) Wait(ctx context.Context) error {
	if r == nil || r.interval <= 0 {
		return nil
	}

	r.mu.Lock()
	scheduled := r.next
	now := time.Now()

	if scheduled.IsZero() || scheduled.Before(now) {
		scheduled = now
	}

	r.next = scheduled.Add(r.interval)
	r.mu.Unlock()

	delay := time.Until(scheduled)
	if delay <= 0 {
		return nil
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
