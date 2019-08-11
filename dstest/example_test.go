package dstest_test

import (
	"context"
	"fmt"
	"time"

	darksky "github.com/twpayne/go-darksky"
	"github.com/twpayne/go-darksky/dstest"
)

func ExampleNewServer() {
	ctx := context.Background()

	s := dstest.NewServer(
		dstest.WithDefaultForecasts(),
	)
	c, err := s.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := c.Forecast(ctx, 34.0219, -118.4814, &darksky.Time{Time: time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC)}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Currently.Icon)

	// Output:
	// partly-cloudy-day
}
