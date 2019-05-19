package darksky

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// A Time is a time that marshals to a UNIX timestamp.
type Time struct {
	time.Time
}

// A ForecastOptions contains options for a forecast request.
type ForecastOptions struct {
	Exclude []string
	Extend  string
	Lang    string
	Units   string
}

// A Alert is a forecast alert.
type Alert struct {
	Description int64    `json:"description,omitempty"`
	Expires     int64    `json:"expires,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity"`
	Time        *Time    `json:"time,omitempty"`
	Title       string   `json:"title,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

// A Currently is the current observation.
type Currently struct {
	ApparentTemperature  float64 `json:"apparentTemperature,omitempty"`
	CloudCover           float64 `json:"cloudCover,omitempty"`
	DewPoint             float64 `json:"dewPoint,omitempty"`
	Humidity             float64 `json:"humidity,omitempty"`
	Icon                 string  `json:"icon,omitempty"`
	NearestStormBearing  float64 `json:"nearestStormBearing,omitempty"`
	NearestStormDistance float64 `json:"nearestStormDistance,omitempty"`
	Ozone                float64 `json:"ozone,omitempty"`
	PrecipIntensity      float64 `json:"precipIntensity,omitempty"`
	PrecipProbability    float64 `json:"precipProbability,omitempty"`
	Pressure             float64 `json:"pressure,omitempty"`
	Summary              string  `json:"summary,omitempty"`
	Temperature          float64 `json:"temperature,omitempty"`
	Time                 *Time   `json:"time,omitempty"`
	UvIndex              float64 `json:"uvIndex,omitempty"`
	Visibility           float64 `json:"visibility,omitempty"`
	WindBearing          float64 `json:"windBearing,omitempty"`
	WindGust             float64 `json:"windGust,omitempty"`
	WindSpeed            float64 `json:"windSpeed,omitempty"`
}

// A DailyData is daily forecast data.
type DailyData struct {
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh,omitempty"`
	ApparentTemperatureHighTime *Time   `json:"apparentTemperatureHighTime,omitempty"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow,omitempty"`
	ApparentTemperatureLowTime  *Time   `json:"apparentTemperatureLowTime,omitempty"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime  *Time   `json:"apparentTemperatureMaxTime,omitempty"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime  float64 `json:"apparentTemperatureMinTime,omitempty"`
	CloudCover                  float64 `json:"cloudCover,omitempty"`
	DewPoint                    float64 `json:"dewPoint,omitempty"`
	Humidity                    float64 `json:"humidity,omitempty"`
	Icon                        string  `json:"icon,omitempty"`
	MoonPhase                   float64 `json:"moonPhase,omitempty"`
	Ozone                       float64 `json:"ozone,omitempty"`
	PrecipIntensity             float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime      *Time   `json:"precipIntensityMaxTime,omitempty"`
	PrecipProbability           float64 `json:"precipProbability,omitempty"`
	PrecipType                  string  `json:"precipType,omitempty"`
	Pressure                    float64 `json:"pressure,omitempty"`
	Summary                     string  `json:"summary,omitempty"`
	SunriseTime                 *Time   `json:"sunriseTime,omitempty"`
	SunsetTime                  *Time   `json:"sunsetTime,omitempty"`
	TemperatureHigh             float64 `json:"temperatureHigh,omitempty"`
	TemperatureHighTime         *Time   `json:"temperatureHighTime,omitempty"`
	TemperatureLow              float64 `json:"temperatureLow,omitempty"`
	TemperatureLowTime          *Time   `json:"temperatureLowTime,omitempty"`
	TemperatureMax              float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime          *Time   `json:"temperatureMaxTime,omitempty"`
	TemperatureMin              float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime          *Time   `json:"temperatureMinTime,omitempty"`
	Time                        *Time   `json:"time,omitempty"`
	UvIndex                     float64 `json:"uvIndex,omitempty"`
	UvIndexTime                 *Time   `json:"uvIndexTime,omitempty"`
	Visibility                  float64 `json:"visibility,omitempty"`
	WindBearing                 float64 `json:"windBearing,omitempty"`
	WindGust                    float64 `json:"windGust,omitempty"`
	WindGustTime                *Time   `json:"windGustTime,omitempty"`
	WindSpeed                   float64 `json:"windSpeed,omitempty"`
}

// A Daily is a daily forecast.
type Daily struct {
	Data    []*DailyData `json:"data,omitempty"`
	Icon    string       `json:"icon,omitempty"`
	Summary string       `json:"summary,omitempty"`
}

// Flags are forecast flags.
type Flags struct {
	DarkSkyUnavailable interface{} `json:"darksky-unavailable,omitempty"`
	NearestStation     float64     `json:"nearest-station,omitempty"`
	Sources            []string    `json:"sources,omitempty"`
	Units              string      `json:"units,omitempty"`
}

// HourlyData is hourly forecast data.
type HourlyData struct {
	ApparentTemperature float64 `json:"apparentTemperature,omitempty"`
	CloudCover          float64 `json:"cloudCover,omitempty"`
	DewPoint            float64 `json:"dewPoint,omitempty"`
	Humidity            float64 `json:"humidity,omitempty"`
	Icon                string  `json:"icon,omitempty"`
	Ozone               float64 `json:"ozone,omitempty"`
	PrecipIntensity     float64 `json:"precipIntensity,omitempty"`
	PrecipProbability   float64 `json:"precipProbability,omitempty"`
	PrecipType          string  `json:"precipType,omitempty"`
	Pressure            float64 `json:"pressure,omitempty"`
	Summary             string  `json:"summary,omitempty"`
	Temperature         float64 `json:"temperature,omitempty"`
	Time                *Time   `json:"time,omitempty"`
	UvIndex             float64 `json:"uvIndex,omitempty"`
	Visibility          float64 `json:"visibility,omitempty"`
	WindBearing         float64 `json:"windBearing,omitempty"`
	WindGust            float64 `json:"windGust,omitempty"`
	WindSpeed           float64 `json:"windSpeed,omitempty"`
}

// A Hourly is an hourly forecast.
type Hourly struct {
	Data    []*HourlyData `json:"data,omitempty"`
	Icon    string        `json:"icon,omitempty"`
	Summary string        `json:"summary,omitempty"`
}

// A MinutelyData is minutely forecast data.
type MinutelyData struct {
	PrecipIntensity      float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
	PrecipProbability    float64 `json:"precipProbability,omitempty"`
	PrecipType           string  `json:"precipType,omitempty"`
	Time                 *Time   `json:"time,omitempty"`
}

// A Minutely is a minutely forecast.
type Minutely struct {
	Data    []*MinutelyData `json:"data,omitempty"`
	Icon    string          `json:"icon,omitempty"`
	Summary string          `json:"summary,omitempty"`
}

// A Forecast is a forecast response.
type Forecast struct {
	Alerts    []*Alert   `json:"alerts,omitempty"`
	Currently *Currently `json:"currently,omitempty"`
	Daily     *Daily     `json:"daily,omitempty"`
	Flags     *Flags     `json:"flags,omitempty"`
	Hourly    *Hourly    `json:"hourly,omitempty"`
	Latitude  float64    `json:"latitude,omitempty"`
	Longitude float64    `json:"longitude,omitempty"`
	Minutely  *Minutely  `json:"minutely,omitempty"`
	Offset    float64    `json:"offset,omitempty"`
	Timezone  string     `json:"timezone,omitempty"`
}

// Forecast returns the forecast for latitude and longitude at time. If time is
// nil or zero then a forecast request is sent. If time is non-nil and non-zero
// then a time machine request is sent.
func (c *Client) Forecast(ctx context.Context, latitude, longitude float64, time *Time, options *ForecastOptions) (*Forecast, error) {
	urlStr := fmt.Sprintf("%s/forecast/%s/%f,%f", c.baseURL, c.key, latitude, longitude)
	if time != nil && !time.IsZero() {
		urlStr += "," + strconv.FormatInt(time.Unix(), 10)
	}
	if options != nil {
		values := url.Values{}
		if len(options.Exclude) != 0 {
			values.Set("exclude", strings.Join(options.Exclude, ","))
		}
		if options.Extend != "" {
			values.Set("extend", options.Extend)
		}
		if options.Lang != "" {
			values.Set("lang", options.Lang)
		}
		if options.Units != "" {
			values.Set("units", options.Units)
		}
		urlStr += "?" + values.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return nil, &ClientError{
			Request:      req,
			Response:     resp,
			ResponseBody: respBody,
		}
	}
	respValue := &Forecast{}
	return respValue, json.NewDecoder(resp.Body).Decode(respValue)
}

// UnmarshalJSON implements the json.Unmarshaler interface. The time is expected
// to be a UNIX timestamp in seconds.
func (t *Time) UnmarshalJSON(data []byte) error {
	var sec int64
	if err := json.Unmarshal(data, &sec); err != nil {
		return err
	}
	t.Time = time.Unix(sec, 0)
	return nil
}
