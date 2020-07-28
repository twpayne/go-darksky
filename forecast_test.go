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
	assert.Equal(t, DefaultUnits, forecast.Flags.Units)
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
	c, err := NewClient(
		WithBaseURL(s.URL),
		WithHTTPClient(s.Client()),
		WithKey("key"),
	)
	require.NoError(t, err)
	_, err = c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
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
		Units:  UnitsSI,
	}
	forecast, err := c.Forecast(context.Background(), 42.3601, -71.0589, nil, options)
	require.NoError(t, err)
	assert.Equal(t, 42.3601, forecast.Latitude)
	assert.Equal(t, -71.0589, forecast.Longitude)
	assert.Equal(t, "America/New_York", forecast.Timezone)
	assert.Equal(t, UnitsSI, forecast.Flags.Units)
	assert.Nil(t, forecast.Alerts)
	assert.Nil(t, forecast.Currently)
	assert.Nil(t, forecast.Daily)
	assert.Nil(t, forecast.Hourly)
	assert.Nil(t, forecast.Minutely)
}

func TestClientTimeMachine(t *testing.T) {
	c := mustNewTestClient(t)
	darkSkyTime := Time{
		Time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	forecast, err := c.Forecast(context.Background(), 42.3601, -71.0589, &darkSkyTime, nil)
	require.NoError(t, err)
	assert.Equal(t, 42.3601, forecast.Latitude)
	assert.Equal(t, -71.0589, forecast.Longitude)
	assert.Equal(t, "America/New_York", forecast.Timezone)
	assert.Equal(t, UnitsUS, forecast.Flags.Units)
	assert.Equal(t, darkSkyTime.Time, forecast.Currently.Time.Time.UTC())
	assert.Nil(t, forecast.Minutely)
	assert.True(t, 24 <= len(forecast.Hourly.Data) && len(forecast.Hourly.Data) < 26)
	assert.Equal(t, 1, len(forecast.Daily.Data))
	assert.Nil(t, forecast.Alerts)
}

func TestClientMetadataCallback(t *testing.T) {
	var lastResponseMetadata *ResponseMetadata
	c := mustNewTestClient(t,
		WithResponseMetadataCallback(func(rm *ResponseMetadata) {
			lastResponseMetadata = rm
		}),
	)
	_, err := c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, lastResponseMetadata)
	assert.Equal(t, http.StatusOK, lastResponseMetadata.StatusCode)
	assert.True(t, lastResponseMetadata.ForecastAPICalls > 0)
	assert.True(t, lastResponseMetadata.ResponseTime > 0)
}

func TestClientInvalidURL(t *testing.T) {
	c, err := NewClient(
		WithBaseURL(""),
		WithKey("%"),
	)
	require.NoError(t, err)
	_, err = c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL escape")
}

func TestClientRequestFail(t *testing.T) {
	c, err := NewClient(
		WithBaseURL("http://0.0.0.0"),
		WithKey("key"),
	)
	require.NoError(t, err)
	_, err = c.Forecast(context.Background(), 42.3601, -71.0589, nil, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}

func TestTimeUnmarshalJSON(t *testing.T) {
	for _, tc := range []struct {
		name         string
		data         []byte
		expectedErr  string
		expectedTime *Time
	}{
		{
			name:         "zero",
			data:         []byte("0"),
			expectedTime: &Time{Time: time.Unix(0, 0)},
		},
		{
			name:        "empty",
			data:        []byte(""),
			expectedErr: "unexpected end of JSON input",
		},
		{
			name:        "string",
			data:        []byte(`""`),
			expectedErr: "json: cannot unmarshal string into Go value of type int64",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			time := &Time{}
			err := time.UnmarshalJSON(tc.data)
			if tc.expectedErr == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedTime, time)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func mustNewTestClient(t *testing.T, options ...ClientOption) *Client {
	key := os.Getenv("DARKSKY_KEY")
	if key == "" {
		t.Fatal("DARKSKY_KEY environment variable not set")
	}
	c, err := NewClient(append([]ClientOption{WithKey(key)}, options...)...)
	require.NoError(t, err)
	return c
}
