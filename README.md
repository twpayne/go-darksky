# go-darksky

[![GoDoc](https://godoc.org/github.com/twpayne/go-darksky?status.svg)](https://godoc.org/github.com/twpayne/go-darksky)
[![Build Status](https://travis-ci.org/twpayne/go-darksky.svg?branch=master)](https://travis-ci.org/twpayne/go-darksky)
[![Coverage Status](https://coveralls.io/repos/github/twpayne/go-darksky/badge.svg)](https://coveralls.io/github/twpayne/go-darksky)

Package `darksky` implements a client for the [Dark Sky weather forecasting
API](https://darksky.net/dev).

## Key features

* Support for all Dark Sky API functionality.
* Idomatic Go API, including support for `context` and Go modules.
* Language matching.
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

	// The forecast varies from day to day. Print something stable.
	fmt.Println(forecast.Timezone)

	// Output:
	// America/New_York
}
```

## Why a new Go Dark Sky client library?

There are several [existing Dark Sky client
libraries](https://darksky.net/dev/docs/libraries). Compared to these, no other
Go library provides all of the following:

* Correct use of types for latitude and longitude: `float64`s, not `string`s.
* Correct use of types for times: `time.Time`s, not `string`s.
* Support for `context`.
* Support for Go modules.
* Rich handling of errors, including both bad requests generic HTTP errors.
* Monitoring hooks.

Adding any of these to an exising Go Dark Sky client library would break API
compatibilty, hence the new client library.

## License

MIT