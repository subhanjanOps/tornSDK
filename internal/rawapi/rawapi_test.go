package rawapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type stubRequester struct {
	t         *testing.T
	wantPath  string
	wantQuery url.Values
	payload   string
}

func (s stubRequester) Do(_ context.Context, req *httpclient.Request, v interface{}) error {
	s.t.Helper()

	if got, want := req.Method, http.MethodGet; got != want {
		s.t.Fatalf("unexpected method: got %q want %q", got, want)
	}

	if got, want := req.Path, s.wantPath; got != want {
		s.t.Fatalf("unexpected path: got %q want %q", got, want)
	}

	if got, want := req.Query.Encode(), s.wantQuery.Encode(); got != want {
		s.t.Fatalf("unexpected query: got %q want %q", got, want)
	}

	return json.Unmarshal([]byte(s.payload), v)
}

func TestNewGetRequest(t *testing.T) {
	query := url.Values{
		"comment": {"bot"},
		"limit":   {"10", "20"},
	}

	req := NewGetRequest("user/events", query)
	if req == nil {
		t.Fatal("expected request")
	}

	if got, want := req.Path, "user/events"; got != want {
		t.Fatalf("unexpected path: got %q want %q", got, want)
	}

	if got, want := req.Query.Get("comment"), "bot"; got != want {
		t.Fatalf("unexpected comment query: got %q want %q", got, want)
	}

	if got, want := len(req.Query["limit"]), 2; got != want {
		t.Fatalf("unexpected limit query count: got %d want %d", got, want)
	}
}

func TestGet(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	response, err := Get(context.Background(), stubRequester{
		t:         t,
		wantPath:  "user/basic",
		wantQuery: query,
		payload:   `{"profile":{"id":123}}`,
	}, "user/basic", query)
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}

	if got, want := string(response), `{"profile":{"id":123}}`; got != want {
		t.Fatalf("unexpected raw response: got %q want %q", got, want)
	}
}
