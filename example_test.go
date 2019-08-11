package darksky_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/twpayne/go-darksky"
)

func ExampleClient_Forecast() {
	c, err := darksky.NewClient(
		darksky.WithKey(os.Getenv("DARKSKY_KEY")),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	forecast, err := c.Forecast(ctx, 42.3601, -71.0589, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// This example requests the current forecast, which varies from day to day.
	// Output only that which is constant.
	fmt.Println(forecast.Latitude)
	fmt.Println(forecast.Longitude)
	fmt.Println(forecast.Timezone)
	fmt.Println(forecast.Flags.Units)

	// Output:
	// 42.3601
	// -71.0589
	// America/New_York
	// us
}

func ExampleClient_Forecast_minimal() {
	c, err := darksky.NewClient(
		darksky.WithKey(os.Getenv("DARKSKY_KEY")),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

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

func ExampleClient_Forecast_timeMachine() {
	c, err := darksky.NewClient(
		darksky.WithKey(os.Getenv("DARKSKY_KEY")),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	t := &darksky.Time{
		Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	forecast, err := c.Forecast(ctx, 42.3601, -71.0589, t, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(forecast.Currently.Icon)

	// Output:
	// cloudy
}
