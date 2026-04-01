package faction

import (
	"context"
	"fmt"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

func (s *Service) GetChainReport(ctx context.Context, chainId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/chainreport", chainId), query)
}

func (s *Service) GetFactionChain(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/chain", id), query)
}

func (s *Service) GetFactionCompletedChains(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/chains", id), query)
}

func (s *Service) GetFactionGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction", query)
}

func (s *Service) GetFactionHoF(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/hof", id), query)
}

func (s *Service) GetFactionLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/lookup", query)
}

func (s *Service) GetFactionMembers(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/members", id), query)
}

func (s *Service) GetFactionRackets(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/rackets", query)
}

func (s *Service) GetFactionRaidsHistory(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/raids", id), query)
}

func (s *Service) GetFactionRankedWarsHistory(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/rankedwars", id), query)
}

func (s *Service) GetFactionSearch(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/search", query)
}

func (s *Service) GetFactionTerritory(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/territory", id), query)
}

func (s *Service) GetFactionTerritoryWarsHistory(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/territorywars", id), query)
}

func (s *Service) GetFactionTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/timestamp", query)
}

func (s *Service) GetFactionWars(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/wars", id), query)
}

func (s *Service) GetMyFactionApplications(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/applications", query)
}

func (s *Service) GetMyFactionAttacks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/attacks", query)
}

func (s *Service) GetMyFactionAttacksSimplified(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/attacksfull", query)
}

func (s *Service) GetMyFactionBalance(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/balance", query)
}

func (s *Service) GetMyFactionChain(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/chain", query)
}

func (s *Service) GetMyFactionCompletedChains(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/chains", query)
}

func (s *Service) GetMyFactionContributors(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/contributors", query)
}

func (s *Service) GetMyFactionHoF(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/hof", query)
}

func (s *Service) GetMyFactionLatestChainReport(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/chainreport", query)
}

func (s *Service) GetMyFactionMembers(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/members", query)
}

func (s *Service) GetMyFactionNews(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/news", query)
}

func (s *Service) GetMyFactionOrganizedCrime(ctx context.Context, crimeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/crime", crimeId), query)
}

func (s *Service) GetMyFactionOrganizedCrimes(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/crimes", query)
}

func (s *Service) GetMyFactionPositions(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/positions", query)
}

func (s *Service) GetMyFactionRaidsHistory(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/raids", query)
}

func (s *Service) GetMyFactionRankedWarsHistory(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/rankedwars", query)
}

func (s *Service) GetMyFactionReports(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/reports", query)
}

func (s *Service) GetMyFactionRevives(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/revives", query)
}

func (s *Service) GetMyFactionRevivesSimplified(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/revivesFull", query)
}

func (s *Service) GetMyFactionStats(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/stats", query)
}

func (s *Service) GetMyFactionTerritory(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/territory", query)
}

func (s *Service) GetMyFactionTerritoryWarsHistory(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/territorywars", query)
}

func (s *Service) GetMyFactionUpgrades(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/upgrades", query)
}

func (s *Service) GetMyFactionWars(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/wars", query)
}

func (s *Service) GetRaidReport(ctx context.Context, raidWarId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/raidreport", raidWarId), query)
}

func (s *Service) GetRankedWarReport(ctx context.Context, rankedWarId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/rankedwarreport", rankedWarId), query)
}

func (s *Service) GetTerritoryOwnership(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/territoryownership", query)
}

func (s *Service) GetTerritoryWarReport(ctx context.Context, territoryWarId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("faction/%s/territorywarreport", territoryWarId), query)
}

func (s *Service) GetWarfare(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "faction/warfare", query)
}
