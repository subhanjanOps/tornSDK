# tornSDK

`tornSDK` is a Go SDK for Torn API v2. It keeps SDK concerns in the SDK: typed models, request transport, error handling, retry behavior, and rate limiting. Game strategy and advisor logic stay outside this repo.

## Current packages

- `client`: public SDK entry point, transport orchestration, retries, rate limiting, and Torn API errors
- `user`: typed core endpoints plus raw wrappers for the remaining user endpoints
- `faction`: typed basic endpoints plus raw wrappers for the remaining faction endpoints
- `forum`: raw wrappers for forum endpoints
- `key`: raw wrappers for key endpoints
- `market`: raw wrappers for market endpoints
- `property`: raw wrappers for property endpoints
- `racing`: raw wrappers for racing endpoints
- `torn`: raw wrappers for torn endpoints

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/subhanjanOps/tornSDK/client"
)

func main() {
	ctx := context.Background()

	sdk := client.New(client.Config{
		APIKey: "YOUR_API_KEY",
	})

	bars, err := sdk.User.GetBars(ctx)
	if err != nil {
		log.Fatal(err)
	}

	basic, err := sdk.User.GetBasic(ctx)
	if err != nil {
		log.Fatal(err)
	}

	profile, err := sdk.User.GetProfile(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s (%d) has %d/%d energy\n", basic.Name, profile.Level, bars.Energy.Current, bars.Energy.Maximum)
}
```

Raw wrapper example:

```go
query := url.Values{"comment": {"my-tool"}}

raw, err := sdk.Torn.GetTornLookup(ctx, query)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(raw))
```

## Configuration

`client.Config` supports:

- `APIKey`
- `BaseURL`
- `HTTPClient`
- `UserAgent`
- `RequestsPerMinute`
- `MaxRetries`
- `RetryWaitMin`
- `RetryWaitMax`

Defaults are aligned to Torn API v2 usage:

- Base URL: `https://api.torn.com/v2`
- Timeout: `15s`
- Rate limit: `100` requests per minute
- Retries: `2` retries with exponential backoff from `1s` to `5s`

Set `RequestsPerMinute` to a negative value to disable SDK throttling. Set `MaxRetries` to a negative value to disable retries.
