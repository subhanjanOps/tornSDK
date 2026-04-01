package client

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/subhanjanOps/tornSDK/faction"
	"github.com/subhanjanOps/tornSDK/internal/httpclient"
	"github.com/subhanjanOps/tornSDK/user"
)

const (
	defaultBaseURL           = "https://api.torn.com/v2"
	defaultTimeout           = 15 * time.Second
	defaultRequestsPerMinute = 100
	defaultMaxRetries        = 2
	defaultRetryWaitMin      = 1 * time.Second
	defaultRetryWaitMax      = 5 * time.Second
	defaultUserAgent         = "tornSDK/0.1"
)

type Config struct {
	APIKey            string
	BaseURL           string
	HTTPClient        *http.Client
	UserAgent         string
	RequestsPerMinute int
	MaxRetries        int
	RetryWaitMin      time.Duration
	RetryWaitMax      time.Duration
}

type Limiter interface {
	Wait(context.Context) error
}

type Client struct {
	apiKey      string
	http        *httpclient.Client
	limiter     Limiter
	retryPolicy RetryPolicy

	User    *user.Service
	Faction *faction.Service
}

func New(config Config) *Client {
	cfg := config.withDefaults()

	c := &Client{
		apiKey:      cfg.APIKey,
		http:        httpclient.New(cfg.BaseURL, cfg.HTTPClient, cfg.UserAgent),
		limiter:     limiterFromConfig(cfg),
		retryPolicy: retryPolicyFromConfig(cfg),
	}

	c.User = user.NewService(c)
	c.Faction = faction.NewService(c)

	return c
}

func (c Config) withDefaults() Config {
	if strings.TrimSpace(c.BaseURL) == "" {
		c.BaseURL = defaultBaseURL
	}

	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: defaultTimeout}
	}

	if strings.TrimSpace(c.UserAgent) == "" {
		c.UserAgent = defaultUserAgent
	}

	if c.RetryWaitMin <= 0 {
		c.RetryWaitMin = defaultRetryWaitMin
	}

	if c.RetryWaitMax <= 0 {
		c.RetryWaitMax = defaultRetryWaitMax
	}

	if c.MaxRetries == 0 {
		c.MaxRetries = defaultMaxRetries
	}

	if c.RequestsPerMinute == 0 {
		c.RequestsPerMinute = defaultRequestsPerMinute
	}

	return c
}

func limiterFromConfig(cfg Config) Limiter {
	if cfg.RequestsPerMinute < 0 {
		return nil
	}

	return NewRateLimiter(cfg.RequestsPerMinute)
}

func retryPolicyFromConfig(cfg Config) RetryPolicy {
	if cfg.MaxRetries < 0 {
		return RetryPolicy{}
	}

	return NewRetryPolicy(cfg.MaxRetries, cfg.RetryWaitMin, cfg.RetryWaitMax)
}
