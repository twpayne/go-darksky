package darksky

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientForecast(t *testing.T) {
	c := mustNewTestClient(t)
	forecast, err := c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 42.3601, forecast.Latitude)
	assert.Equal(t, -71.0589, forecast.Longitude)
	assert.Equal(t, "America/New_York", forecast.Timezone)
	assert.Equal(t, "us", forecast.Flags.Units)
	assert.True(t, time.Since(forecast.Currently.Time.Time) < 10*time.Second)
}

func TestClientTimeMachine(t *testing.T) {
	c := mustNewTestClient(t)
	darkSkyTime := Time{Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)}
	forecast, err := c.Forecast(context.Background(), 42.3601, -71.0589, &darkSkyTime, nil)
	require.NoError(t, err)
	assert.Equal(t, 42.3601, forecast.Latitude)
	assert.Equal(t, -71.0589, forecast.Longitude)
	assert.Equal(t, "America/New_York", forecast.Timezone)
	assert.Equal(t, "us", forecast.Flags.Units)
	assert.Equal(t, darkSkyTime.Time, forecast.Currently.Time.Time.UTC())
	assert.Nil(t, forecast.Minutely)
	assert.Equal(t, 24, len(forecast.Hourly.Data))
	assert.Equal(t, 1, len(forecast.Daily.Data))
	assert.Nil(t, forecast.Alerts)
}

func mustNewTestClient(t *testing.T) *Client {
	key := os.Getenv("DARKSKY_KEY")
	if key == "" {
		t.Fatal("DARKSKY_KEY environment variable not set")
	}
	return NewClient(WithKey(key))
}
