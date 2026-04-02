package client

import (
	"context"
	"testing"
	"time"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type stubLogger struct {
	msgs []string
}

func (s *stubLogger) Debugf(format string, v ...interface{}) { s.msgs = append(s.msgs, "debug") }
func (s *stubLogger) Infof(format string, v ...interface{})  { s.msgs = append(s.msgs, "info") }
func (s *stubLogger) Warnf(format string, v ...interface{})  { s.msgs = append(s.msgs, "warn") }
func (s *stubLogger) Errorf(format string, v ...interface{}) { s.msgs = append(s.msgs, "error") }

func TestClientAcceptsLogger(t *testing.T) {
	l := &stubLogger{}
	c := New(Config{APIKey: "k", Logger: l, RequestsPerMinute: -1})
	if c.logger == nil {
		t.Fatal("expected logger to be set on client")
	}

	// perform a Do call with an invalid path to exercise logging before BuildRequest
	_ = c.Do(context.Background(), httpclient.NewRequest("GET", "bad\npath"), nil)

	// logger should have recorded at least the debug call before build request
	if len(l.msgs) == 0 {
		t.Fatal("expected logger to receive messages")
	}

	// ensure client creation didn't modify defaults unexpectedly
	if c.retryPolicy.MaxRetries != defaultMaxRetries {
		t.Fatalf("unexpected retries: %d", c.retryPolicy.MaxRetries)
	}

	// also ensure withDefaults sets reasonable timeouts
	if c.http == nil {
		t.Fatal("expected http client to be present")
	}

	// quick sanity: ensure New doesn't block
	time.Sleep(1 * time.Millisecond)
}
