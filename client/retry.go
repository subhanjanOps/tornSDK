package client

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

type RetryPolicy struct {
	MaxRetries int
	MinWait    time.Duration
	MaxWait    time.Duration
}

func NewRetryPolicy(maxRetries int, minWait, maxWait time.Duration) RetryPolicy {
	if maxRetries < 0 {
		maxRetries = 0
	}

	if minWait <= 0 {
		minWait = defaultRetryWaitMin
	}

	if maxWait <= 0 {
		maxWait = defaultRetryWaitMax
	}

	if maxWait < minWait {
		maxWait = minWait
	}

	return RetryPolicy{
		MaxRetries: maxRetries,
		MinWait:    minWait,
		MaxWait:    maxWait,
	}
}

func (p RetryPolicy) NextBackoff(attempt int, resp *http.Response, err error) (time.Duration, bool) {
	if attempt >= p.MaxRetries || !IsTemporary(err) {
		return 0, false
	}

	if resp != nil {
		if retryAfter, ok := parseRetryAfter(resp.Header.Get("Retry-After")); ok {
			return retryAfter, true
		}
	}

	return p.backoff(attempt), true
}

func (p RetryPolicy) backoff(attempt int) time.Duration {
	wait := p.MinWait
	for i := 0; i < attempt; i++ {
		wait *= 2
		if wait >= p.MaxWait {
			return p.MaxWait
		}
	}

	if wait > p.MaxWait {
		return p.MaxWait
	}

	return wait
}

func parseRetryAfter(value string) (time.Duration, bool) {
	if value == "" {
		return 0, false
	}

	if seconds, err := strconv.Atoi(value); err == nil {
		if seconds < 0 {
			seconds = 0
		}

		return time.Duration(seconds) * time.Second, true
	}

	when, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}

	wait := time.Until(when)
	if wait < 0 {
		wait = 0
	}

	return wait, true
}

func sleepContext(ctx context.Context, wait time.Duration) error {
	if wait <= 0 {
		return nil
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
