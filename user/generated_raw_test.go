package user

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

func TestUserRawMethods(t *testing.T) {
	query := url.Values{"comment": {"sdk"}}

	cases := []struct {
		name     string
		wantPath string
		call     func(*Service) (rawapi.Response, error)
	}{
		{
			name:     "GetMyAmmo",
			wantPath: "user/ammo",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyAmmo(context.Background(), query) },
		},
		{
			name:     "GetMyAttacks",
			wantPath: "user/attacks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyAttacks(context.Background(), query) },
		},
		{
			name:     "GetMyAttacksSimplified",
			wantPath: "user/attacksfull",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyAttacksSimplified(context.Background(), query)
			},
		},
		{
			name:     "GetMyAvailableOrganizedCrimes",
			wantPath: "user/organizedcrimes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyAvailableOrganizedCrimes(context.Background(), query)
			},
		},
		{
			name:     "GetMyBounties",
			wantPath: "user/bounties",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyBounties(context.Background(), query) },
		},
		{
			name:     "GetMyCalendarTime",
			wantPath: "user/calendar",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyCalendarTime(context.Background(), query) },
		},
		{
			name:     "GetMyCompetitionInfo",
			wantPath: "user/competition",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyCompetitionInfo(context.Background(), query) },
		},
		{
			name:     "GetMyContactsList",
			wantPath: "user/list",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyContactsList(context.Background(), query) },
		},
		{
			name:     "GetMyCooldowns",
			wantPath: "user/cooldowns",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyCooldowns(context.Background(), query) },
		},
		{
			name:     "GetMyCrimes",
			wantPath: "user/crimeId-value/crimes",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyCrimes(context.Background(), "crimeId-value", query)
			},
		},
		{
			name:     "GetMyDetailedTrade",
			wantPath: "user/tradeId-value/trade",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyDetailedTrade(context.Background(), "tradeId-value", query)
			},
		},
		{
			name:     "GetMyDiscord",
			wantPath: "user/discord",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyDiscord(context.Background(), query) },
		},
		{
			name:     "GetMyEducation",
			wantPath: "user/education",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyEducation(context.Background(), query) },
		},
		{
			name:     "GetMyEnlistedCars",
			wantPath: "user/enlistedcars",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyEnlistedCars(context.Background(), query) },
		},
		{
			name:     "GetMyEquipment",
			wantPath: "user/equipment",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyEquipment(context.Background(), query) },
		},
		{
			name:     "GetMyEvents",
			wantPath: "user/events",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyEvents(context.Background(), query) },
		},
		{
			name:     "GetMyFaction",
			wantPath: "user/faction",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyFaction(context.Background(), query) },
		},
		{
			name:     "GetMyForumFeed",
			wantPath: "user/forumfeed",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyForumFeed(context.Background(), query) },
		},
		{
			name:     "GetMyForumFriendsUpdates",
			wantPath: "user/forumfriends",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyForumFriendsUpdates(context.Background(), query)
			},
		},
		{
			name:     "GetMyForumPosts",
			wantPath: "user/forumposts",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyForumPosts(context.Background(), query) },
		},
		{
			name:     "GetMyForumSubscribedThreads",
			wantPath: "user/forumsubscribedthreads",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyForumSubscribedThreads(context.Background(), query)
			},
		},
		{
			name:     "GetMyForumThreads",
			wantPath: "user/forumthreads",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyForumThreads(context.Background(), query) },
		},
		{
			name:     "GetMyHoF",
			wantPath: "user/hof",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyHoF(context.Background(), query) },
		},
		{
			name:     "GetMyHonors",
			wantPath: "user/honors",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyHonors(context.Background(), query) },
		},
		{
			name:     "GetMyIcons",
			wantPath: "user/icons",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyIcons(context.Background(), query) },
		},
		{
			name:     "GetMyItemMarketListings",
			wantPath: "user/itemmarket",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyItemMarketListings(context.Background(), query)
			},
		},
		{
			name:     "GetMyJob",
			wantPath: "user/job",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyJob(context.Background(), query) },
		},
		{
			name:     "GetMyJobPoints",
			wantPath: "user/jobpoints",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyJobPoints(context.Background(), query) },
		},
		{
			name:     "GetMyJobRanks",
			wantPath: "user/jobranks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyJobRanks(context.Background(), query) },
		},
		{
			name:     "GetMyLogs",
			wantPath: "user/log",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyLogs(context.Background(), query) },
		},
		{
			name:     "GetMyMedals",
			wantPath: "user/medals",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyMedals(context.Background(), query) },
		},
		{
			name:     "GetMyMerits",
			wantPath: "user/merits",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyMerits(context.Background(), query) },
		},
		{
			name:     "GetMyMessages",
			wantPath: "user/messages",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyMessages(context.Background(), query) },
		},
		{
			name:     "GetMyMissions",
			wantPath: "user/missions",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyMissions(context.Background(), query) },
		},
		{
			name:     "GetMyMoney",
			wantPath: "user/money",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyMoney(context.Background(), query) },
		},
		{
			name:     "GetMyNewEvents",
			wantPath: "user/newevents",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyNewEvents(context.Background(), query) },
		},
		{
			name:     "GetMyNewMessages",
			wantPath: "user/newmessages",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyNewMessages(context.Background(), query) },
		},
		{
			name:     "GetMyNotifications",
			wantPath: "user/notifications",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyNotifications(context.Background(), query) },
		},
		{
			name:     "GetMyOrganizedCrime",
			wantPath: "user/organizedcrime",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyOrganizedCrime(context.Background(), query) },
		},
		{
			name:     "GetMyPersonalStats",
			wantPath: "user/personalstats",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyPersonalStats(context.Background(), query) },
		},
		{
			name:     "GetMyProperties",
			wantPath: "user/properties",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyProperties(context.Background(), query) },
		},
		{
			name:     "GetMyProperty",
			wantPath: "user/property",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyProperty(context.Background(), query) },
		},
		{
			name:     "GetMyRaces",
			wantPath: "user/races",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyRaces(context.Background(), query) },
		},
		{
			name:     "GetMyRacingRecords",
			wantPath: "user/racingrecords",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyRacingRecords(context.Background(), query) },
		},
		{
			name:     "GetMyRefills",
			wantPath: "user/refills",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyRefills(context.Background(), query) },
		},
		{
			name:     "GetMyReports",
			wantPath: "user/reports",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyReports(context.Background(), query) },
		},
		{
			name:     "GetMyRevives",
			wantPath: "user/revives",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyRevives(context.Background(), query) },
		},
		{
			name:     "GetMyRevivesSimplified",
			wantPath: "user/revivesFull",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyRevivesSimplified(context.Background(), query)
			},
		},
		{
			name:     "GetMySkills",
			wantPath: "user/skills",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMySkills(context.Background(), query) },
		},
		{
			name:     "GetMyStocks",
			wantPath: "user/stocks",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyStocks(context.Background(), query) },
		},
		{
			name:     "GetMyTrades",
			wantPath: "user/trades",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyTrades(context.Background(), query) },
		},
		{
			name:     "GetMyTravelInformation",
			wantPath: "user/travel",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyTravelInformation(context.Background(), query)
			},
		},
		{
			name:     "GetMyVirusCodingInformation",
			wantPath: "user/virus",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetMyVirusCodingInformation(context.Background(), query)
			},
		},
		{
			name:     "GetMyWeaponExp",
			wantPath: "user/weaponexp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyWeaponExp(context.Background(), query) },
		},
		{
			name:     "GetMyWorkstats",
			wantPath: "user/workstats",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetMyWorkstats(context.Background(), query) },
		},
		{
			name:     "GetUserBasicInformation",
			wantPath: "user/id-value/basic",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserBasicInformation(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserBounties",
			wantPath: "user/id-value/bounties",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserBounties(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserCompetitionInfo",
			wantPath: "user/id-value/competition",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserCompetitionInfo(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserDiscord",
			wantPath: "user/id-value/discord",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserDiscord(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserFaction",
			wantPath: "user/id-value/faction",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserFaction(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserForumPosts",
			wantPath: "user/id-value/forumposts",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserForumPosts(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserForumThreads",
			wantPath: "user/id-value/forumthreads",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserForumThreads(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserGeneric",
			wantPath: "user",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetUserGeneric(context.Background(), query) },
		},
		{
			name:     "GetUserHoF",
			wantPath: "user/id-value/hof",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserHoF(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserIcons",
			wantPath: "user/id-value/icons",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserIcons(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserJob",
			wantPath: "user/id-value/job",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserJob(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserLookup",
			wantPath: "user/lookup",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetUserLookup(context.Background(), query) },
		},
		{
			name:     "GetUserPersonalStats",
			wantPath: "user/id-value/personalstats",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserPersonalStats(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserProfile",
			wantPath: "user/id-value/profile",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserProfile(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserProperties",
			wantPath: "user/id-value/properties",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserProperties(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserProperty",
			wantPath: "user/id-value/property",
			call: func(s *Service) (rawapi.Response, error) {
				return s.GetUserProperty(context.Background(), "id-value", query)
			},
		},
		{
			name:     "GetUserTimestamp",
			wantPath: "user/timestamp",
			call:     func(s *Service) (rawapi.Response, error) { return s.GetUserTimestamp(context.Background(), query) },
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
