package market

import (
	"context"
	"encoding/json"
	"net/url"
	"testing"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type stubRequester struct {
	t        *testing.T
	pages    [][]byte
	calls    int
	wantPath string
}

func (s *stubRequester) Do(_ context.Context, req *httpclient.Request, v interface{}) error {
	s.calls++
	if got, want := req.Path, s.wantPath; got != want {
		s.t.Fatalf("unexpected path: got %q want %q", got, want)
	}

	// return next page
	idx := s.calls - 1
	if idx < 0 || idx >= len(s.pages) {
		s.t.Fatalf("unexpected page index %d", idx)
	}

	return json.Unmarshal(s.pages[idx], v)
}

func TestPagerIteratesPages(t *testing.T) {
	pages := [][]byte{
		[]byte(`[{"id":1},{"id":2}]`),
		[]byte(`[{"id":3}]`),
	}

	req := &stubRequester{t: t, pages: pages, wantPath: "market/auctionhouse"}

	pager := NewPager(req, "market/auctionhouse", url.Values{"q": {"v"}}, 2)

	ctx := context.Background()

	// first page
	resp, done, err := pager.Next(ctx)
	if err != nil || done {
		t.Fatalf("unexpected first Next: err=%v done=%v", err, done)
	}
	if string(resp) != string(pages[0]) {
		t.Fatalf("unexpected first page: %s", resp)
	}

	// second page
	resp, done, err = pager.Next(ctx)
	if err != nil || !done {
		t.Fatalf("unexpected second Next: err=%v done=%v", err, done)
	}
	if string(resp) != string(pages[1]) {
		t.Fatalf("unexpected second page: %s", resp)
	}
}
