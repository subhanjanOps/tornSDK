package torn

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

func (s *Service) GetSpecificTornStock(ctx context.Context, stockId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/stocks", stockId), query)
}

func (s *Service) GetTornAttackLog(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/attacklog", query)
}

func (s *Service) GetTornBounties(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/bounties", query)
}

func (s *Service) GetTornCalendar(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/calendar", query)
}

func (s *Service) GetTornCrimes(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/crimes", query)
}

func (s *Service) GetTornEducation(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/education", query)
}

func (s *Service) GetTornElimination(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/elimination", query)
}

func (s *Service) GetTornEliminationTeam(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/eliminationteam", id), query)
}

func (s *Service) GetTornFactionHoF(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/factionhof", query)
}

func (s *Service) GetTornFactionTree(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/factiontree", query)
}

func (s *Service) GetTornGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn", query)
}

func (s *Service) GetTornHoF(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/hof", query)
}

func (s *Service) GetTornHonors(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/honors", query)
}

func (s *Service) GetTornHonorsSpecific(ctx context.Context, ids string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/honors", ids), query)
}

func (s *Service) GetTornItemAmmo(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/itemammo", query)
}

func (s *Service) GetTornItemDetails(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/itemdetails", id), query)
}

func (s *Service) GetTornItemMods(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/itemmods", query)
}

func (s *Service) GetTornItems(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/items", query)
}

func (s *Service) GetTornItemsSpecific(ctx context.Context, ids string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/items", ids), query)
}

func (s *Service) GetTornLogCategories(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/logcategories", query)
}

func (s *Service) GetTornLogTypes(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/logtypes", query)
}

func (s *Service) GetTornLogTypesSpecific(ctx context.Context, logCategoryId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/logtypes", logCategoryId), query)
}

func (s *Service) GetTornLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/lookup", query)
}

func (s *Service) GetTornMedals(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/medals", query)
}

func (s *Service) GetTornMedalsSpecific(ctx context.Context, ids string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/medals", ids), query)
}

func (s *Service) GetTornMerits(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/merits", query)
}

func (s *Service) GetTornOrganizedCrimes(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/organizedcrimes", query)
}

func (s *Service) GetTornProperties(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/properties", query)
}

func (s *Service) GetTornStocks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/stocks", query)
}

func (s *Service) GetTornSubcrimes(ctx context.Context, crimeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("torn/%s/subcrimes", crimeId), query)
}

func (s *Service) GetTornTerritory(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/territory", query)
}

func (s *Service) GetTornTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "torn/timestamp", query)
}
