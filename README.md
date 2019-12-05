# go-darksky

[![GoDoc](https://godoc.org/github.com/twpayne/go-darksky?status.svg)](https://godoc.org/github.com/twpayne/go-darksky)
[![Coverage Status](https://coveralls.io/repos/github/twpayne/go-darksky/badge.svg)](https://coveralls.io/github/twpayne/go-darksky)

Package `darksky` implements a client for the [Dark Sky weather forecasting
API](https://darksky.net/dev).

## Key features

* Support for all Dark Sky API functionality.
* Idomatic Go API, including support for `context` and Go modules.
* Language matching.
* Rich handling of errors, including both bad requests generic HTTP errors.
* Fully tested, including error conditions.
* Mock client for offline testing.
* Monitoring hooks.

## Example

```go
func ExampleClient_Forecast() {
	c := darksky.NewClient(
		darksky.WithKey(os.Getenv("DARKSKY_KEY")),
	)

	ctx := context.Background()
	forecast, err := c.Forecast(ctx, 42.3601, -71.0589, nil, &darksky.ForecastOptions{
		Units: darksky.UnitsSI,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(forecast)
}
```

## License

MIT
