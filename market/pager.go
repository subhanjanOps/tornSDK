package market

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/subhanjanOps/tornSDK/internal/rawapi"
)

// Pager iterates paginated market endpoints using `limit` and `offset` query params.
// It is conservative: it stops when a page returns an empty array or when the
// returned JSON has fewer items than the requested page size.
type Pager struct {
	client rawapi.Requester
	path   string
	base   url.Values
	limit  int
	offset int
	done   bool
}

func NewPager(client rawapi.Requester, path string, baseQuery url.Values, pageSize int) *Pager {
	if baseQuery == nil {
		baseQuery = make(url.Values)
	}

	return &Pager{
		client: client,
		path:   path,
		base:   baseQuery,
		limit:  pageSize,
		offset: 0,
	}
}

// Next returns the next page as raw JSON. When no more pages are available,
// it returns an empty response and nil error with done==true.
func (p *Pager) Next(ctx context.Context) (rawapi.Response, bool, error) {
	if p.done {
		return nil, true, nil
	}

	q := url.Values{}
	for k, v := range p.base {
		q[k] = append([]string(nil), v...)
	}

	if p.limit > 0 {
		q.Set("limit", intToString(p.limit))
	}
	q.Set("offset", intToString(p.offset))

	resp, err := rawapi.Get(ctx, p.client, p.path, q)
	if err != nil {
		return nil, false, err
	}

	// determine if resp is an array and its length
	var arr []json.RawMessage
	if err := json.Unmarshal(resp, &arr); err != nil {
		// not an array — return as single page and mark done
		p.done = true
		return resp, true, nil
	}

	if len(arr) == 0 || (p.limit > 0 && len(arr) < p.limit) {
		p.done = true
	}

	p.offset += p.limit
	return resp, p.done, nil
}

func intToString(i int) string {
	if i <= 0 {
		return "0"
	}
	return strconv.Itoa(i)
}
