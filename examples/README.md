Examples

This folder contains small example programs that demonstrate how to use the SDK.

pager: demonstrates using `market.Pager` to iterate paginated market endpoints.

Run the pager example:

```bash
go run ./examples/pager
```

Note: The example disables the rate limiter by default for local runs. Set a
valid `APIKey` in the example or pass one via environment/config when using the
real Torn API.
