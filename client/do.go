package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/subhanjanOps/tornSDK/internal/httpclient"
)

func (c *Client) Do(ctx context.Context, req *Request, v interface{}) error {
	if c == nil {
		return errors.New("nil client")
	}

	if req == nil {
		return errors.New("nil request")
	}

	prepared := req.Clone()
	prepared.SetQuery("key", c.apiKey)

	for attempt := 0; ; attempt++ {
		if c.limiter != nil {
			if err := c.limiter.Wait(ctx); err != nil {
				return err
			}
		}

		httpReq, err := c.http.BuildRequest(ctx, prepared)
		if err != nil {
			return err
		}

		resp, err := c.http.Do(httpReq)
		if err == nil {
			err = decodeResponse(resp, v)
			if err == nil {
				return nil
			}
		}

		wait, shouldRetry := c.retryPolicy.NextBackoff(attempt, resp, err)
		if !shouldRetry {
			return err
		}

		if err := sleepContext(ctx, wait); err != nil {
			return err
		}
	}
}

func decodeResponse(resp *http.Response, v interface{}) error {
	if resp == nil {
		return errors.New("nil response")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if apiErr := parseAPIError(body); apiErr != nil {
		return apiErr
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       shortenBody(body),
		}
	}

	if v == nil || len(bytes.TrimSpace(body)) == 0 {
		return nil
	}

	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}

	return nil
}

var _ interface {
	Do(context.Context, *httpclient.Request, interface{}) error
} = (*Client)(nil)
