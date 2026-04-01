package faction

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

func TestFactionRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetChainReport",
			wantPath: "faction/chainId-value/chainreport",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetChainReport(context.Background(), "chainId-value", query)
			},
		},
		{
			name:     "GetFactionChain",
			wantPath: "faction/id-value/chain",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionChain(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionCompletedChains",
			wantPath: "faction/id-value/chains",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionCompletedChains(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionGeneric",
			wantPath: "faction",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetFactionGeneric(context.Background(), query) },
		},
		{
			name:     "GetFactionHoF",
			wantPath: "faction/id-value/hof",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionHoF(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionLookup",
			wantPath: "faction/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetFactionLookup(context.Background(), query) },
		},
		{
			name:     "GetFactionMembers",
			wantPath: "faction/id-value/members",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionMembers(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionRackets",
			wantPath: "faction/rackets",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetFactionRackets(context.Background(), query) },
		},
		{
			name:     "GetFactionRaidsHistory",
			wantPath: "faction/id-value/raids",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionRaidsHistory(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionRankedWarsHistory",
			wantPath: "faction/id-value/rankedwars",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionRankedWarsHistory(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionSearch",
			wantPath: "faction/search",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetFactionSearch(context.Background(), query) },
		},
		{
			name:     "GetFactionTerritory",
			wantPath: "faction/id-value/territory",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionTerritory(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionTerritoryWarsHistory",
			wantPath: "faction/id-value/territorywars",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionTerritoryWarsHistory(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetFactionTimestamp",
			wantPath: "faction/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetFactionTimestamp(context.Background(), query) },
		},
		{
			name:     "GetFactionWars",
			wantPath: "faction/id-value/wars",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetFactionWars(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetMyFactionApplications",
			wantPath: "faction/applications",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionApplications(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionAttacks",
			wantPath: "faction/attacks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionAttacks(context.Background(), query) },
		},
		{
			name:     "GetMyFactionAttacksSimplified",
			wantPath: "faction/attacksfull",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionAttacksSimplified(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionBalance",
			wantPath: "faction/balance",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionBalance(context.Background(), query) },
		},
		{
			name:     "GetMyFactionChain",
			wantPath: "faction/chain",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionChain(context.Background(), query) },
		},
		{
			name:     "GetMyFactionCompletedChains",
			wantPath: "faction/chains",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionCompletedChains(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionContributors",
			wantPath: "faction/contributors",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionContributors(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionHoF",
			wantPath: "faction/hof",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionHoF(context.Background(), query) },
		},
		{
			name:     "GetMyFactionLatestChainReport",
			wantPath: "faction/chainreport",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionLatestChainReport(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionMembers",
			wantPath: "faction/members",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionMembers(context.Background(), query) },
		},
		{
			name:     "GetMyFactionNews",
			wantPath: "faction/news",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionNews(context.Background(), query) },
		},
		{
			name:     "GetMyFactionOrganizedCrime",
			wantPath: "faction/crimeId-value/crime",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionOrganizedCrime(context.Background(), "crimeId-value", query)
			},
		},
		{
			name:     "GetMyFactionOrganizedCrimes",
			wantPath: "faction/crimes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionOrganizedCrimes(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionPositions",
			wantPath: "faction/positions",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionPositions(context.Background(), query) },
		},
		{
			name:     "GetMyFactionRaidsHistory",
			wantPath: "faction/raids",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionRaidsHistory(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionRankedWarsHistory",
			wantPath: "faction/rankedwars",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionRankedWarsHistory(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionReports",
			wantPath: "faction/reports",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionReports(context.Background(), query) },
		},
		{
			name:     "GetMyFactionRevives",
			wantPath: "faction/revives",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionRevives(context.Background(), query) },
		},
		{
			name:     "GetMyFactionRevivesSimplified",
			wantPath: "faction/revivesFull",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionRevivesSimplified(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionStats",
			wantPath: "faction/stats",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionStats(context.Background(), query) },
		},
		{
			name:     "GetMyFactionTerritory",
			wantPath: "faction/territory",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionTerritory(context.Background(), query) },
		},
		{
			name:     "GetMyFactionTerritoryWarsHistory",
			wantPath: "faction/territorywars",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyFactionTerritoryWarsHistory(context.Background(), query)
			},
		},
		{
			name:     "GetMyFactionUpgrades",
			wantPath: "faction/upgrades",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionUpgrades(context.Background(), query) },
		},
		{
			name:     "GetMyFactionWars",
			wantPath: "faction/wars",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFactionWars(context.Background(), query) },
		},
		{
			name:     "GetRaidReport",
			wantPath: "faction/raidWarId-value/raidreport",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetRaidReport(context.Background(), "raidWarId-value", query)
			},
		},
		{
			name:     "GetRankedWarReport",
			wantPath: "faction/rankedWarId-value/rankedwarreport",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetRankedWarReport(context.Background(), "rankedWarId-value", query)
			},
		},
		{
			name:     "GetTerritoryOwnership",
			wantPath: "faction/territoryownership",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetTerritoryOwnership(context.Background(), query) },
		},
		{
			name:     "GetTerritoryWarReport",
			wantPath: "faction/territoryWarId-value/territorywarreport",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetTerritoryWarReport(context.Background(), "territoryWarId-value", query)
			},
		},
		{
			name:     "GetWarfare",
			wantPath: "faction/warfare",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetWarfare(context.Background(), query) },
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
