package key

import (
	"context"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

type Service struct {
	client rawapi.Requester
}

func NewService(c rawapi.Requester) *Service {
	return &Service{client: c}
}

func (s *Service) GetKeyGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "key", query)
}

func (s *Service) GetKeyInfo(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "key/info", query)
}

func (s *Service) GetKeyLog(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "key/log", query)
}
