package client

import (
	"net/http"
	"testing"
	"time"
)

func TestConfigWithDefaults(t *testing.T) {
	cfg := (Config{}).withDefaults()

	if got, want := cfg.BaseURL, defaultBaseURL; got != want {
		t.Fatalf("unexpected base URL: got %q want %q", got, want)
	}

	if cfg.HTTPClient == nil {
		t.Fatal("expected default HTTP client")
	}

	if got, want := cfg.HTTPClient.Timeout, defaultTimeout; got != want {
		t.Fatalf("unexpected timeout: got %s want %s", got, want)
	}

	if got, want := cfg.UserAgent, defaultUserAgent; got != want {
		t.Fatalf("unexpected user agent: got %q want %q", got, want)
	}

	if got, want := cfg.RetryWaitMin, defaultRetryWaitMin; got != want {
		t.Fatalf("unexpected retry wait min: got %s want %s", got, want)
	}

	if got, want := cfg.RetryWaitMax, defaultRetryWaitMax; got != want {
		t.Fatalf("unexpected retry wait max: got %s want %s", got, want)
	}

	if got, want := cfg.MaxRetries, defaultMaxRetries; got != want {
		t.Fatalf("unexpected max retries: got %d want %d", got, want)
	}

	if got, want := cfg.RequestsPerMinute, defaultRequestsPerMinute; got != want {
		t.Fatalf("unexpected RPM: got %d want %d", got, want)
	}
}

func TestConfigWithDefaultsPreservesExplicitValues(t *testing.T) {
	httpClient := &http.Client{Timeout: 3 * time.Second}
	cfg := (Config{
		BaseURL:           "https://example.com/custom",
		HTTPClient:        httpClient,
		UserAgent:         "custom-agent",
		RequestsPerMinute: -1,
		MaxRetries:        -1,
		RetryWaitMin:      2 * time.Second,
		RetryWaitMax:      4 * time.Second,
	}).withDefaults()

	if got, want := cfg.BaseURL, "https://example.com/custom"; got != want {
		t.Fatalf("unexpected base URL: got %q want %q", got, want)
	}

	if cfg.HTTPClient != httpClient {
		t.Fatal("expected explicit HTTP client to be preserved")
	}

	if got, want := cfg.UserAgent, "custom-agent"; got != want {
		t.Fatalf("unexpected user agent: got %q want %q", got, want)
	}

	if got, want := cfg.RequestsPerMinute, -1; got != want {
		t.Fatalf("unexpected RPM: got %d want %d", got, want)
	}

	if got, want := cfg.MaxRetries, -1; got != want {
		t.Fatalf("unexpected max retries: got %d want %d", got, want)
	}

	if got, want := cfg.RetryWaitMin, 2*time.Second; got != want {
		t.Fatalf("unexpected retry wait min: got %s want %s", got, want)
	}

	if got, want := cfg.RetryWaitMax, 4*time.Second; got != want {
		t.Fatalf("unexpected retry wait max: got %s want %s", got, want)
	}
}

func TestNewInitializesServicesAndPolicies(t *testing.T) {
	c := New(Config{
		APIKey:            "secret",
		RequestsPerMinute: -1,
		MaxRetries:        -1,
	})

	if c == nil {
		t.Fatal("expected client")
	}

	if c.User == nil {
		t.Fatal("expected user service")
	}

	if c.Faction == nil {
		t.Fatal("expected faction service")
	}

	if c.limiter != nil {
		t.Fatal("expected limiter to be disabled for negative RPM")
	}

	if got := c.retryPolicy; got != (RetryPolicy{}) {
		t.Fatalf("unexpected retry policy: %+v", got)
	}
}

func TestLimiterAndRetryPolicyFromConfig(t *testing.T) {
	if limiter := limiterFromConfig(Config{RequestsPerMinute: -1}); limiter != nil {
		t.Fatal("expected nil limiter for negative RPM")
	}

	limiter := limiterFromConfig(Config{RequestsPerMinute: 60})
	rateLimiter, ok := limiter.(*RateLimiter)
	if !ok || rateLimiter == nil {
		t.Fatal("expected rate limiter")
	}

	if got, want := rateLimiter.interval, time.Second; got != want {
		t.Fatalf("unexpected interval: got %s want %s", got, want)
	}

	if got := retryPolicyFromConfig(Config{MaxRetries: -1}); got != (RetryPolicy{}) {
		t.Fatalf("unexpected disabled retry policy: %+v", got)
	}

	got := retryPolicyFromConfig(Config{
		MaxRetries:   2,
		RetryWaitMin: time.Second,
		RetryWaitMax: 3 * time.Second,
	})
	want := NewRetryPolicy(2, time.Second, 3*time.Second)
	if got != want {
		t.Fatalf("unexpected retry policy: got %+v want %+v", got, want)
	}
}
