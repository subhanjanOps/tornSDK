package forum

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

type rawStubRequester struct {
	t         *testing.T
	wantPath  string
	wantQuery url.Values
	payload   string
}

func (s rawStubRequester) Do(_ context.Context, req *httpclient.Request, v interface{}) error {
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

func TestForumRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetForumAllThreads",
			wantPath: "forum/threads",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetForumAllThreads(context.Background(), query) },
		},
		{
			name:     "GetForumCategories",
			wantPath: "forum/categories",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetForumCategories(context.Background(), query) },
		},
		{
			name:     "GetForumGeneric",
			wantPath: "forum",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetForumGeneric(context.Background(), query) },
		},
		{
			name:     "GetForumLookup",
			wantPath: "forum/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetForumLookup(context.Background(), query) },
		},
		{
			name:     "GetForumThread",
			wantPath: "forum/threadId-value/thread",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetForumThread(context.Background(), "threadId-value", query)
			},
		},
		{
			name:     "GetForumThreadPosts",
			wantPath: "forum/threadId-value/posts",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetForumThreadPosts(context.Background(), "threadId-value", query)
			},
		},
		{
			name:     "GetForumThreads",
			wantPath: "forum/categoryIds-value/threads",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetForumThreads(context.Background(), "categoryIds-value", query)
			},
		},
		{
			name:     "GetForumTimestamp",
			wantPath: "forum/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetForumTimestamp(context.Background(), query) },
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(rawStubRequester{
				t:         t,
				wantPath:  tc.wantPath,
				wantQuery: query,
				payload:   `{"ok":true}`,
			})

			response, err := tc.call(service)
			if err != nil {
				t.Fatalf("method returned error: %v", err)
			}

			if got, want := string(response), `{"ok":true}`; got != want {
				t.Fatalf("unexpected response: got %q want %q", got, want)
			}
		})
	}
}
