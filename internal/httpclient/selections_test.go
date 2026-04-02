package httpclient

import (
	"testing"
)

func TestSelectionsBuilder(t *testing.T) {
	s := NewSelections("bars", "profile", "  ")
	if got, want := s.String(), "bars,profile"; got != want {
		t.Fatalf("unexpected selections string: got %q want %q", got, want)
	}

	req := NewRequest("GET", "user/bars")
	s.Apply(req)
	if got := req.Query.Get("selections"); got != "bars,profile" {
		t.Fatalf("unexpected request selections: got %q", got)
	}

	// empty selections should remove the query
	s2 := NewSelections()
	s2.Apply(req)
	if got := req.Query.Get("selections"); got != "" {
		t.Fatalf("expected empty selections to remove query, got %q", got)
	}
}
