package property

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

func (s *Service) GetProperty(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("property/%s/property", id), query)
}

func (s *Service) GetPropertyGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "property", query)
}

func (s *Service) GetPropertyLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "property/lookup", query)
}

func (s *Service) GetPropertyTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "property/timestamp", query)
}
