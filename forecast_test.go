package darksky

import (
	"context"
	"net/http"
	"net/http/httptest"
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

func TestClientBadRequest(t *testing.T) {
	c := mustNewTestClient(t)
	_, err := c.Forecast(context.Background(), -180, -90, nil, nil)
	require.Error(t, err)
	e, ok := err.(*Error)
	require.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, e.Details.Code)
	assert.Equal(t, "The given location is invalid.", e.Error())
}

func TestClientInternalServerError(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("response body"))
	}))
	c := NewClient(
		WithBaseURL(s.URL),
		WithHTTPClient(s.Client()),
		WithKey("key"),
	)
	_, err := c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
	require.Error(t, err)
	e, ok := err.(*Error)
	require.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, e.Response.StatusCode)
	assert.Equal(t, s.URL+"/forecast/key/42.360100,-71.058900: 500 Internal Server Error: response body", e.Error())
}

func TestClientForecastOptions(t *testing.T) {
	c := mustNewTestClient(t)
	options := &ForecastOptions{
		Exclude: []Block{
			BlockAlerts,
			BlockCurrently,
			BlockDaily,
			BlockHourly,
			BlockMinutely,
		},
		Extend: ExtendHourly,
		Lang:   LangAR,
		Units:  "si",
	}
	forecast, err := c.Forecast(context.Background(), 42.3601, -71.0589, nil, options)
	require.NoError(t, err)
	assert.Equal(t, 42.3601, forecast.Latitude)
	assert.Equal(t, -71.0589, forecast.Longitude)
	assert.Equal(t, "America/New_York", forecast.Timezone)
	assert.Equal(t, "si", forecast.Flags.Units)
	assert.Nil(t, forecast.Alerts)
	assert.Nil(t, forecast.Currently)
	assert.Nil(t, forecast.Daily)
	assert.Nil(t, forecast.Hourly)
	assert.Nil(t, forecast.Minutely)
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
	return NewClient(
		WithKey(key),
	)
}
