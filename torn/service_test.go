package torn

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

func TestTornRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetSpecificTornStock",
			wantPath: "torn/stockId-value/stocks",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetSpecificTornStock(context.Background(), "stockId-value", query)
			},
		},
		{
			name:     "GetTornAttackLog",
			wantPath: "torn/attacklog",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornAttackLog(context.Background(), query) },
		},
		{
			name:     "GetTornBounties",
			wantPath: "torn/bounties",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornBounties(context.Background(), query) },
		},
		{
			name:     "GetTornCalendar",
			wantPath: "torn/calendar",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornCalendar(context.Background(), query) },
		},
		{
			name:     "GetTornCrimes",
			wantPath: "torn/crimes",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornCrimes(context.Background(), query) },
		},
		{
			name:     "GetTornEducation",
			wantPath: "torn/education",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornEducation(context.Background(), query) },
		},
		{
			name:     "GetTornElimination",
			wantPath: "torn/elimination",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornElimination(context.Background(), query) },
		},
		{
			name:     "GetTornEliminationTeam",
			wantPath: "torn/id-value/eliminationteam",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornEliminationTeam(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetTornFactionHoF",
			wantPath: "torn/factionhof",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornFactionHoF(context.Background(), query) },
		},
		{
			name:     "GetTornFactionTree",
			wantPath: "torn/factiontree",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornFactionTree(context.Background(), query) },
		},
		{
			name:     "GetTornGeneric",
			wantPath: "torn",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornGeneric(context.Background(), query) },
		},
		{
			name:     "GetTornHoF",
			wantPath: "torn/hof",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornHoF(context.Background(), query) },
		},
		{
			name:     "GetTornHonors",
			wantPath: "torn/honors",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornHonors(context.Background(), query) },
		},
		{
			name:     "GetTornHonorsSpecific",
			wantPath: "torn/ids-value/honors",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornHonorsSpecific(context.Background(), "ids-value", query)
			},
		},
		{
			name:     "GetTornItemAmmo",
			wantPath: "torn/itemammo",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornItemAmmo(context.Background(), query) },
		},
		{
			name:     "GetTornItemDetails",
			wantPath: "torn/id-value/itemdetails",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornItemDetails(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetTornItemMods",
			wantPath: "torn/itemmods",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornItemMods(context.Background(), query) },
		},
		{
			name:     "GetTornItems",
			wantPath: "torn/items",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornItems(context.Background(), query) },
		},
		{
			name:     "GetTornItemsSpecific",
			wantPath: "torn/ids-value/items",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornItemsSpecific(context.Background(), "ids-value", query)
			},
		},
		{
			name:     "GetTornLogCategories",
			wantPath: "torn/logcategories",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornLogCategories(context.Background(), query) },
		},
		{
			name:     "GetTornLogTypes",
			wantPath: "torn/logtypes",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornLogTypes(context.Background(), query) },
		},
		{
			name:     "GetTornLogTypesSpecific",
			wantPath: "torn/logCategoryId-value/logtypes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornLogTypesSpecific(context.Background(), "logCategoryId-value", query)
			},
		},
		{
			name:     "GetTornLookup",
			wantPath: "torn/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornLookup(context.Background(), query) },
		},
		{
			name:     "GetTornMedals",
			wantPath: "torn/medals",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornMedals(context.Background(), query) },
		},
		{
			name:     "GetTornMedalsSpecific",
			wantPath: "torn/ids-value/medals",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornMedalsSpecific(context.Background(), "ids-value", query)
			},
		},
		{
			name:     "GetTornMerits",
			wantPath: "torn/merits",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornMerits(context.Background(), query) },
		},
		{
			name:     "GetTornOrganizedCrimes",
			wantPath: "torn/organizedcrimes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornOrganizedCrimes(context.Background(), query)
			},
		},
		{
			name:     "GetTornProperties",
			wantPath: "torn/properties",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornProperties(context.Background(), query) },
		},
		{
			name:     "GetTornStocks",
			wantPath: "torn/stocks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornStocks(context.Background(), query) },
		},
		{
			name:     "GetTornSubcrimes",
			wantPath: "torn/crimeId-value/subcrimes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTornSubcrimes(context.Background(), "crimeId-value", query)
			},
		},
		{
			name:     "GetTornTerritory",
			wantPath: "torn/territory",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornTerritory(context.Background(), query) },
		},
		{
			name:     "GetTornTimestamp",
			wantPath: "torn/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTornTimestamp(context.Background(), query) },
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
