package client

import "github.com/subhanjanOps/tornSDK/internal/httpclient"

type Request = httpclient.Request

func NewRequest(method, path string) *Request {
	return httpclient.NewRequest(method, path)
}
