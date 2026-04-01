package racing

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

func TestRacingRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetRacingCarUpgrades",
			wantPath: "racing/carupgrades",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingCarUpgrades(context.Background(), query) },
		},
		{
			name:     "GetRacingCars",
			wantPath: "racing/cars",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingCars(context.Background(), query) },
		},
		{
			name:     "GetRacingGeneric",
			wantPath: "racing",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingGeneric(context.Background(), query) },
		},
		{
			name:     "GetRacingLookup",
			wantPath: "racing/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingLookup(context.Background(), query) },
		},
		{
			name:     "GetRacingRaceDetails",
			wantPath: "racing/raceId-value/race",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetRacingRaceDetails(context.Background(), "raceId-value", query)
			},
		},
		{
			name:     "GetRacingRaces",
			wantPath: "racing/races",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingRaces(context.Background(), query) },
		},
		{
			name:     "GetRacingTimestamp",
			wantPath: "racing/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingTimestamp(context.Background(), query) },
		},
		{
			name:     "GetRacingTrackRecords",
			wantPath: "racing/trackId-value/records",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetRacingTrackRecords(context.Background(), "trackId-value", query)
			},
		},
		{
			name:     "GetRacingTracks",
			wantPath: "racing/tracks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetRacingTracks(context.Background(), query) },
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
