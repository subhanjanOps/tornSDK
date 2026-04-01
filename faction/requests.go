package faction

import (
	"fmt"
	"net/http"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

const resourcePath = "faction/"

const basicPath = resourcePath + "basic"

func newOwnerFactionRequest() *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, basicPath)
}

func newFactionByIDRequest(id int) *httpclient.Request {
	return httpclient.NewRequest(http.MethodGet, fmt.Sprintf("%s%d/basic", resourcePath, id))
}
