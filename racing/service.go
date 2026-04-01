package racing

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

func (s *Service) GetRacingCarUpgrades(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/carupgrades", query)
}

func (s *Service) GetRacingCars(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/cars", query)
}

func (s *Service) GetRacingGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing", query)
}

func (s *Service) GetRacingLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/lookup", query)
}

func (s *Service) GetRacingRaceDetails(ctx context.Context, raceId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("racing/%s/race", raceId), query)
}

func (s *Service) GetRacingRaces(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/races", query)
}

func (s *Service) GetRacingTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/timestamp", query)
}

func (s *Service) GetRacingTrackRecords(ctx context.Context, trackId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("racing/%s/records", trackId), query)
}

func (s *Service) GetRacingTracks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "racing/tracks", query)
}
