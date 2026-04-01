package rawapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

type Response = json.RawMessage

type Requester interface {
	Do(context.Context, *httpclient.Request, interface{}) error
}

func NewGetRequest(path string, query url.Values) *httpclient.Request {
	req := httpclient.NewRequest(http.MethodGet, path)
	for key, values := range query {
		for _, value := range values {
			req.AddQuery(key, value)
		}
	}

	return req
}

func Get(ctx context.Context, requester Requester, path string, query url.Values) (Response, error) {
	var response Response
	err := requester.Do(ctx, NewGetRequest(path, query), &response)
	return response, err
}
