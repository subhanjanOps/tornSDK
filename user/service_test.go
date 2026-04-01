package user

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

func TestGetBasicUsesV2ProfileEnvelope(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   basicPath,
		payload:    `{"profile":{"id":123,"name":"Subha","level":42,"gender":"Male","status":{"description":"Okay","details":null,"state":"Okay","color":"green","until":null}}}`,
	})

	basic, err := service.GetBasic(context.Background())
	if err != nil {
		t.Fatalf("GetBasic returned error: %v", err)
	}

	if got, want := basic.ID, 123; got != want {
		t.Fatalf("unexpected ID: got %d want %d", got, want)
	}

	if got, want := basic.Name, "Subha"; got != want {
		t.Fatalf("unexpected name: got %q want %q", got, want)
	}
}

func TestGetBarsUsesV2BarsEnvelope(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   barsPath,
		payload:    `{"bars":{"energy":{"current":10,"maximum":150},"nerve":{"current":5,"maximum":60},"happy":{"current":2500,"maximum":5000},"life":{"current":1000,"maximum":1000},"chain":null}}`,
	})

	bars, err := service.GetBars(context.Background())
	if err != nil {
		t.Fatalf("GetBars returned error: %v", err)
	}

	if got, want := bars.Energy.Current, 10; got != want {
		t.Fatalf("unexpected energy current: got %d want %d", got, want)
	}

	if got, want := bars.Life.Maximum, 1000; got != want {
		t.Fatalf("unexpected life maximum: got %d want %d", got, want)
	}
}

func TestGetProfileUsesV2ProfilePath(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   profilePath,
		payload:    `{"profile":{"id":123,"name":"Subha","level":42,"rank":"Star","title":"Beginner","age":365,"signed_up":1700000000,"faction_id":null,"honor_id":7,"property":{"id":9,"name":"Shack"},"image":null,"gender":"Male","revivable":true,"role":"Civilian","status":{"description":"Okay","details":null,"state":"Okay","color":"green","until":null},"spouse":null,"awards":5,"friends":10,"enemies":2,"forum_posts":3,"karma":4,"last_action":{"status":"Offline","timestamp":1700000001,"relative":"1 minute ago"},"life":{"current":500,"maximum":600},"donator_status":null}}`,
	})

	profile, err := service.GetProfile(context.Background())
	if err != nil {
		t.Fatalf("GetProfile returned error: %v", err)
	}

	if got, want := profile.Property.Name, "Shack"; got != want {
		t.Fatalf("unexpected property name: got %q want %q", got, want)
	}

	if got, want := profile.Life.Current, 500; got != want {
		t.Fatalf("unexpected life current: got %d want %d", got, want)
	}
}

func TestGetBattleStatsUsesV2BattleStatsEnvelope(t *testing.T) {
	service := NewService(stubRequester{
		t:          t,
		wantMethod: http.MethodGet,
		wantPath:   battleStatsPath,
		payload:    `{"battlestats":{"strength":{"value":1,"modifier":0,"modifiers":[]},"defense":{"value":2,"modifier":0,"modifiers":[]},"speed":{"value":3,"modifier":0,"modifiers":[]},"dexterity":{"value":4,"modifier":0,"modifiers":[]},"total":10}}`,
	})

	stats, err := service.GetBattleStats(context.Background())
	if err != nil {
		t.Fatalf("GetBattleStats returned error: %v", err)
	}

	if got, want := stats.Strength.Value, int64(1); got != want {
		t.Fatalf("unexpected strength value: got %d want %d", got, want)
	}

	if got, want := stats.Total, int64(10); got != want {
		t.Fatalf("unexpected total: got %d want %d", got, want)
	}
}
