package market

import (
	"context"
	"fmt"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

type Service struct {
	client rawapi.Requester
}

func NewService(c rawapi.Requester) *Service {
	return &Service{client: c}
}

func (s *Service) GetMarketAuctionHouse(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "market/auctionhouse", query)
}

func (s *Service) GetMarketAuctionHouseItem(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/auctionhouse", id), query)
}

func (s *Service) GetMarketAuctionHouseListing(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/auctionhouselisting", id), query)
}

func (s *Service) GetMarketBazaar(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "market/bazaar", query)
}

func (s *Service) GetMarketBazaarItem(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/bazaar", id), query)
}

func (s *Service) GetMarketGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "market", query)
}

func (s *Service) GetMarketItemMarketItem(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/itemmarket", id), query)
}

func (s *Service) GetMarketLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "market/lookup", query)
}

func (s *Service) GetMarketProperties(ctx context.Context, propertyTypeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/properties", propertyTypeId), query)
}

func (s *Service) GetMarketPropertiesRental(ctx context.Context, propertyTypeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("market/%s/rentals", propertyTypeId), query)
}

func (s *Service) GetMarketTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "market/timestamp", query)
}
