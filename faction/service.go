package faction

import (
	"context"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type requester interface {
	Do(context.Context, *httpclient.Request, interface{}) error
}

type Service struct {
	client requester
}

func NewService(c requester) *Service {
	return &Service{client: c}
}

func (s *Service) GetBasic(ctx context.Context) (*Basic, error) {
	var response basicResponse
	err := s.client.Do(ctx, newOwnerFactionRequest(), &response)
	return &response.Basic, err
}

func (s *Service) GetBasicByID(ctx context.Context, id int) (*Basic, error) {
	var response basicResponse
	err := s.client.Do(ctx, newFactionByIDRequest(id), &response)
	return &response.Basic, err
}
