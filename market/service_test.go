package market

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

func TestMarketRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetMarketAuctionHouse",
			wantPath: "market/auctionhouse",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMarketAuctionHouse(context.Background(), query) },
		},
		{
			name:     "GetMarketAuctionHouseItem",
			wantPath: "market/id-value/auctionhouse",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketAuctionHouseItem(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetMarketAuctionHouseListing",
			wantPath: "market/id-value/auctionhouselisting",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketAuctionHouseListing(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetMarketBazaar",
			wantPath: "market/bazaar",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMarketBazaar(context.Background(), query) },
		},
		{
			name:     "GetMarketBazaarItem",
			wantPath: "market/id-value/bazaar",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketBazaarItem(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetMarketGeneric",
			wantPath: "market",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMarketGeneric(context.Background(), query) },
		},
		{
			name:     "GetMarketItemMarketItem",
			wantPath: "market/id-value/itemmarket",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketItemMarketItem(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetMarketLookup",
			wantPath: "market/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMarketLookup(context.Background(), query) },
		},
		{
			name:     "GetMarketProperties",
			wantPath: "market/propertyTypeId-value/properties",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketProperties(context.Background(), "propertyTypeId-value", query)
			},
		},
		{
			name:     "GetMarketPropertiesRental",
			wantPath: "market/propertyTypeId-value/rentals",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMarketPropertiesRental(context.Background(), "propertyTypeId-value", query)
			},
		},
		{
			name:     "GetMarketTimestamp",
			wantPath: "market/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMarketTimestamp(context.Background(), query) },
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
