package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Query   url.Values
	Headers http.Header
}

func NewRequest(method, path string) *Request {
	return &Request{
		Method:  method,
		Path:    strings.TrimLeft(path, "/"),
		Query:   make(url.Values),
		Headers: make(http.Header),
	}
}

func (r *Request) Clone() *Request {
	if r == nil {
		return nil
	}

	query := make(url.Values, len(r.Query))
	for key, values := range r.Query {
		query[key] = append([]string(nil), values...)
	}

	headers := make(http.Header, len(r.Headers))
	for key, values := range r.Headers {
		headers[key] = append([]string(nil), values...)
	}

	return &Request{
		Method:  r.Method,
		Path:    r.Path,
		Query:   query,
		Headers: headers,
	}
}

func (r *Request) AddQuery(key string, values ...string) *Request {
	if r == nil {
		return nil
	}

	if r.Query == nil {
		r.Query = make(url.Values)
	}

	for _, value := range values {
		r.Query.Add(key, value)
	}

	return r
}

func (r *Request) SetQuery(key string, values ...string) *Request {
	if r == nil {
		return nil
	}

	if r.Query == nil {
		r.Query = make(url.Values)
	}

	if len(values) == 0 {
		delete(r.Query, key)
		return r
	}

	r.Query.Del(key)
	for _, value := range values {
		r.Query.Add(key, value)
	}

	return r
}

func (r *Request) SetSelections(selections ...string) *Request {
	filtered := make([]string, 0, len(selections))
	for _, selection := range selections {
		if trimmed := strings.TrimSpace(selection); trimmed != "" {
			filtered = append(filtered, trimmed)
		}
	}

	if len(filtered) == 0 {
		return r.SetQuery("selections")
	}

	return r.SetQuery("selections", strings.Join(filtered, ","))
}

type Client struct {
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

func New(baseURL string, httpClient *http.Client, userAgent string) *Client {
	return &Client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		userAgent:  userAgent,
	}
}

func (c *Client) BuildRequest(ctx context.Context, req *Request) (*http.Request, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}

	method := req.Method
	if method == "" {
		method = http.MethodGet
	}

	fullURL := c.baseURL + "/" + strings.TrimLeft(req.Path, "/")
	if len(req.Query) > 0 {
		fullURL += "?" + req.Query.Encode()
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, fullURL, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		httpReq.Header.Set("User-Agent", c.userAgent)
	}

	for key, values := range req.Headers {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	return httpReq, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
