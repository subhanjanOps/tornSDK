package user

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
	err := s.client.Do(ctx, newBasicRequest(), &response)
	return &response.Profile, err
}

func (s *Service) GetProfile(ctx context.Context) (*Profile, error) {
	var response profileResponse
	err := s.client.Do(ctx, newProfileRequest(), &response)
	return &response.Profile, err
}

func (s *Service) GetBars(ctx context.Context) (*Bars, error) {
	var response barsResponse
	err := s.client.Do(ctx, newBarsRequest(), &response)
	return &response.Bars, err
}

func (s *Service) GetBattleStats(ctx context.Context) (*BattleStats, error) {
	var response battleStatsResponse
	err := s.client.Do(ctx, newBattleStatsRequest(), &response)
	return &response.BattleStats, err
}
