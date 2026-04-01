package faction

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type stubRequester struct {
	t          *testing.T
	wantMethod string
	wantPath   string
	payload    string
}

func (s stubRequester) Do(_ context.Context, req *httpclient.Request, v interface{}) error {
	s.t.Helper()

	if got, want := req.Method, s.wantMethod; got != want {
		s.t.Fatalf("unexpected method: got %q want %q", got, want)
	}

	if got, want := req.Path, s.wantPath; got != want {
		s.t.Fatalf("unexpected path: got %q want %q", got, want)
	}

	return json.Unmarshal([]byte(s.payload), v)
}

func TestGetBasicUsesV2BasicEnvelope(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   basicPath,
		payload:    `{"basic":{"id":321,"name":"Faction","tag":"TAG","tag_image":"img","leader_id":1,"co_leader_id":2,"respect":1000,"days_old":50,"capacity":100,"members":80,"is_enlisted":true,"rank":{"level":1,"name":"Bronze","division":2,"position":3,"wins":4},"best_chain":500}}`,
	})

	basic, err := service.GetBasic(context.Background())
	if err != nil {
		t.Fatalf("GetBasic returned error: %v", err)
	}

	if got, want := basic.ID, 321; got != want {
		t.Fatalf("unexpected ID: got %d want %d", got, want)
	}

	if got, want := basic.Rank.Name, "Bronze"; got != want {
		t.Fatalf("unexpected rank name: got %q want %q", got, want)
	}
}

func TestGetBasicByIDUsesFactionPath(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   "faction/999/basic",
		payload:    `{"basic":{"id":999,"name":"Faction","tag":"TAG","tag_image":"img","leader_id":1,"co_leader_id":2,"respect":1000,"days_old":50,"capacity":100,"members":80,"is_enlisted":null,"rank":{"level":1,"name":"Bronze","division":2,"position":3,"wins":4},"best_chain":500}}`,
	})

	basic, err := service.GetBasicByID(context.Background(), 999)
	if err != nil {
		t.Fatalf("GetBasicByID returned error: %v", err)
	}

	if got, want := basic.ID, 999; got != want {
		t.Fatalf("unexpected ID: got %d want %d", got, want)
	}
}
