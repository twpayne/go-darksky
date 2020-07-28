package dstest_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/twpayne/go-darksky"
	"github.com/twpayne/go-darksky/dstest"
)

func TestServer(t *testing.T) {
	for _, tc := range []struct {
		name               string
		request            dstest.Request
		expectedStatusCode int
		testForecast       func(*testing.T, *darksky.Forecast)
	}{
		{
			name: "not_found",
			request: dstest.Request{
				Latitude:  0,
				Longitude: 0,
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.DefaultUnits,
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "latitude_out_of_range",
			request: dstest.Request{
				Latitude:  -100,
				Longitude: 0,
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.DefaultUnits,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "longitude_out_of_range",
			request: dstest.Request{
				Latitude:  0,
				Longitude: 190,
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.DefaultUnits,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "santamonica_20190501",
			request: dstest.Request{
				Latitude:  34.0219,
				Longitude: -118.4814,
				Time:      darksky.Time{Time: time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC)},
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.DefaultUnits,
			},
			expectedStatusCode: http.StatusOK,
			testForecast: func(t *testing.T, f *darksky.Forecast) {
				assert.Equal(t, "America/Los_Angeles", f.Timezone)
			},
		},
		{
			name: "santamonica_hourly_si",
			request: dstest.Request{
				Latitude:  34.0219,
				Longitude: -118.4814,
				Extend:    darksky.ExtendHourly,
				Lang:      darksky.DefaultLang,
				Units:     darksky.UnitsSI,
			},
			expectedStatusCode: http.StatusOK,
			testForecast: func(t *testing.T, f *darksky.Forecast) {
				assert.Equal(t, 169, len(f.Hourly.Data))
			},
		},
		{
			name: "santamonica_exclude_si",
			request: dstest.Request{
				Latitude:  34.0219,
				Longitude: -118.4814,
				Exclude:   "alerts,currently,daily,flags,minutely",
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.UnitsSI,
			},
			expectedStatusCode: http.StatusOK,
			testForecast: func(t *testing.T, f *darksky.Forecast) {
				assert.Nil(t, f.Alerts)
				assert.Nil(t, f.Currently)
				assert.Nil(t, f.Daily)
				assert.Nil(t, f.Flags)
				assert.Nil(t, f.Minutely)
			},
		},
		{
			name: "santamonica_exclude_out_of_order_si",
			request: dstest.Request{
				Latitude:  34.0219,
				Longitude: -118.4814,
				Exclude:   "minutely,flags,daily,currently,alerts",
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.DefaultLang,
				Units:     darksky.UnitsSI,
			},
			expectedStatusCode: http.StatusOK,
			testForecast: func(t *testing.T, f *darksky.Forecast) {
				assert.Nil(t, f.Alerts)
				assert.Nil(t, f.Currently)
				assert.Nil(t, f.Daily)
				assert.Nil(t, f.Flags)
				assert.Nil(t, f.Minutely)
			},
		},
		{
			name: "santamonica_fr",
			request: dstest.Request{
				Latitude:  34.0219,
				Longitude: -118.4814,
				Extend:    darksky.DefaultExtend,
				Lang:      darksky.LangFR,
				Units:     darksky.DefaultUnits,
			},
			expectedStatusCode: http.StatusOK,
			testForecast: func(t *testing.T, f *darksky.Forecast) {
				assert.Equal(t, "Ciel Dégagé", f.Currently.Summary)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := dstest.NewServer(
				dstest.WithDefaultForecasts(),
			)
			c, err := s.NewClient()
			require.NoError(t, err)
			f, err := c.Forecast(context.Background(), tc.request.Latitude, tc.request.Longitude, &tc.request.Time, tc.request.Options())
			if tc.expectedStatusCode != http.StatusOK {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			if tc.testForecast != nil {
				tc.testForecast(t, f)
			}
		})
	}
}
