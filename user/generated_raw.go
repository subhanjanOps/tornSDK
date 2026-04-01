package user

import (
	"context"
	"fmt"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

func (s *Service) GetMyAmmo(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/ammo", query)
}

func (s *Service) GetMyAttacks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/attacks", query)
}

func (s *Service) GetMyAttacksSimplified(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/attacksfull", query)
}

func (s *Service) GetMyAvailableOrganizedCrimes(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/organizedcrimes", query)
}

func (s *Service) GetMyBounties(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/bounties", query)
}

func (s *Service) GetMyCalendarTime(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/calendar", query)
}

func (s *Service) GetMyCompetitionInfo(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/competition", query)
}

func (s *Service) GetMyContactsList(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/list", query)
}

func (s *Service) GetMyCooldowns(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/cooldowns", query)
}

func (s *Service) GetMyCrimes(ctx context.Context, crimeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/crimes", crimeId), query)
}

func (s *Service) GetMyDetailedTrade(ctx context.Context, tradeId string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/trade", tradeId), query)
}

func (s *Service) GetMyDiscord(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/discord", query)
}

func (s *Service) GetMyEducation(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/education", query)
}

func (s *Service) GetMyEnlistedCars(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/enlistedcars", query)
}

func (s *Service) GetMyEquipment(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/equipment", query)
}

func (s *Service) GetMyEvents(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/events", query)
}

func (s *Service) GetMyFaction(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/faction", query)
}

func (s *Service) GetMyForumFeed(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/forumfeed", query)
}

func (s *Service) GetMyForumFriendsUpdates(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/forumfriends", query)
}

func (s *Service) GetMyForumPosts(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/forumposts", query)
}

func (s *Service) GetMyForumSubscribedThreads(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/forumsubscribedthreads", query)
}

func (s *Service) GetMyForumThreads(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/forumthreads", query)
}

func (s *Service) GetMyHoF(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/hof", query)
}

func (s *Service) GetMyHonors(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/honors", query)
}

func (s *Service) GetMyIcons(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/icons", query)
}

func (s *Service) GetMyItemMarketListings(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/itemmarket", query)
}

func (s *Service) GetMyJob(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/job", query)
}

func (s *Service) GetMyJobPoints(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/jobpoints", query)
}

func (s *Service) GetMyJobRanks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/jobranks", query)
}

func (s *Service) GetMyLogs(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/log", query)
}

func (s *Service) GetMyMedals(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/medals", query)
}

func (s *Service) GetMyMerits(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/merits", query)
}

func (s *Service) GetMyMessages(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/messages", query)
}

func (s *Service) GetMyMissions(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/missions", query)
}

func (s *Service) GetMyMoney(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/money", query)
}

func (s *Service) GetMyNewEvents(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/newevents", query)
}

func (s *Service) GetMyNewMessages(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/newmessages", query)
}

func (s *Service) GetMyNotifications(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/notifications", query)
}

func (s *Service) GetMyOrganizedCrime(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/organizedcrime", query)
}

func (s *Service) GetMyPersonalStats(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/personalstats", query)
}

func (s *Service) GetMyProperties(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/properties", query)
}

func (s *Service) GetMyProperty(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/property", query)
}

func (s *Service) GetMyRaces(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/races", query)
}

func (s *Service) GetMyRacingRecords(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/racingrecords", query)
}

func (s *Service) GetMyRefills(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/refills", query)
}

func (s *Service) GetMyReports(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/reports", query)
}

func (s *Service) GetMyRevives(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/revives", query)
}

func (s *Service) GetMyRevivesSimplified(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/revivesFull", query)
}

func (s *Service) GetMySkills(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/skills", query)
}

func (s *Service) GetMyStocks(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/stocks", query)
}

func (s *Service) GetMyTrades(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/trades", query)
}

func (s *Service) GetMyTravelInformation(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/travel", query)
}

func (s *Service) GetMyVirusCodingInformation(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/virus", query)
}

func (s *Service) GetMyWeaponExp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/weaponexp", query)
}

func (s *Service) GetMyWorkstats(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/workstats", query)
}

func (s *Service) GetUserBasicInformation(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/basic", id), query)
}

func (s *Service) GetUserBounties(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/bounties", id), query)
}

func (s *Service) GetUserCompetitionInfo(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/competition", id), query)
}

func (s *Service) GetUserDiscord(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/discord", id), query)
}

func (s *Service) GetUserFaction(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/faction", id), query)
}

func (s *Service) GetUserForumPosts(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/forumposts", id), query)
}

func (s *Service) GetUserForumThreads(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/forumthreads", id), query)
}

func (s *Service) GetUserGeneric(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user", query)
}

func (s *Service) GetUserHoF(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/hof", id), query)
}

func (s *Service) GetUserIcons(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/icons", id), query)
}

func (s *Service) GetUserJob(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/job", id), query)
}

func (s *Service) GetUserLookup(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/lookup", query)
}

func (s *Service) GetUserPersonalStats(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/personalstats", id), query)
}

func (s *Service) GetUserProfile(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/profile", id), query)
}

func (s *Service) GetUserProperties(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/properties", id), query)
}

func (s *Service) GetUserProperty(ctx context.Context, id string, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, fmt.Sprintf("user/%s/property", id), query)
}

func (s *Service) GetUserTimestamp(ctx context.Context, query url.Values) (rawapi.Response, error) {
	return rawapi.Get(ctx, s.client, "user/timestamp", query)
}
