package user

import (
	"net/http"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

const (
	barsPath        = "user/bars"
	basicPath       = "user/basic"
	battleStatsPath = "user/battlestats"
	profilePath     = "user/profile"
)

func newBarsRequest() *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, barsPath)
}

func newBasicRequest() *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, basicPath)
}

func newBattleStatsRequest() *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, battleStatsPath)
}

func newProfileRequest() *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, profilePath)
}
