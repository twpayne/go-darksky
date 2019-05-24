package darksky

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// A Block is a data block in a response.
type Block string

// Blocks.
const (
	BlockAlerts    Block = "alerts"
	BlockCurrently Block = "currently"
	BlockDaily     Block = "daily"
	BlockFlags     Block = "flags"
	BlockHourly    Block = "hourly"
	BlockMinutely  Block = "minutely"
)

// An Extend is what can be extended.
type Extend string

// Extends.
const (
	ExtendHourly = "hourly"
	ExtendNone   = ""
)

// A Lang is a language.
type Lang string

// Langs.
const (
	LangAR        Lang = "ar"          // Arabic
	LangAZ        Lang = "az"          // Azerbaijani
	LangBE        Lang = "be"          // Belarusian
	LangBG        Lang = "bg"          // Bulgarian
	LangBN        Lang = "bn"          // Bengali
	LangBS        Lang = "bs"          // Bosnian
	LangCA        Lang = "ca"          // Catalan
	LangCS        Lang = "cs"          // Czech
	LangDA        Lang = "da"          // Danish
	LangDE        Lang = "de"          // German
	LangEL        Lang = "el"          // Greek
	LangEN        Lang = "en"          // English (which is the default)
	LangEO        Lang = "eo"          // Esperanto
	LangES        Lang = "es"          // Spanish
	LangET        Lang = "et"          // Estonian
	LangFI        Lang = "fi"          // Finnish
	LangFR        Lang = "fr"          // French
	LangHE        Lang = "he"          // Hebrew
	LangHI        Lang = "hi"          // Hindi
	LangHR        Lang = "hr"          // Croatian
	LangHU        Lang = "hu"          // Hungarian
	LangID        Lang = "id"          // Indonesian
	LangIS        Lang = "is"          // Icelandic
	LangIT        Lang = "it"          // Italian
	LangJA        Lang = "ja"          // Japanese
	LangKA        Lang = "ka"          // Georgian
	LangKN        Lang = "kn"          // Kannada
	LangKO        Lang = "ko"          // Korean
	LangKW        Lang = "kw"          // Cornish
	LangLV        Lang = "lv"          // Latvian
	LangML        Lang = "ml"          // Malayam
	LangMR        Lang = "mr"          // Marathi
	LangNB        Lang = "nb"          // Norwegian Bokmål
	LangNL        Lang = "nl"          // Dutch
	LangNO        Lang = "no"          // Norwegian Bokmål (alias for nb)
	LangPA        Lang = "pa"          // Punjabi
	LangPL        Lang = "pl"          // Polish
	LangPT        Lang = "pt"          // Portuguese
	LangRO        Lang = "ro"          // Romanian
	LangRU        Lang = "ru"          // Russian
	LangSK        Lang = "sk"          // Slovak
	LangSL        Lang = "sl"          // Slovenian
	LangSR        Lang = "sr"          // Serbian
	LangSV        Lang = "sv"          // Swedish
	LangTA        Lang = "ta"          // Tamil
	LangTE        Lang = "te"          // Telugu
	LangTET       Lang = "tet"         // Tetum
	LangTR        Lang = "tr"          // Turkish
	LangUK        Lang = "uk"          // Ukrainian
	LangUR        Lang = "ur"          // Urdu
	LangXPigLatin Lang = "x-pig-latin" // Igpay Atinlay
	LangZH        Lang = "zh"          // simplified Chinese
	LangZHTW      Lang = "zh-tw"       // traditional ChineLang
)

// A Units is a system of units.
type Units string

// Units.
const (
	UnitsAuto Units = "auto"
	UnitsCA   Units = "ca"
	UnitsSI   Units = "si"
	UnitsUK2  Units = "uk2"
	UnitsUS   Units = "us"
)

// A Time is a time that unmarshals from a UNIX timestamp.
type Time struct {
	time.Time
}

// A ForecastOptions contains options for a forecast request.
type ForecastOptions struct {
	Exclude []Block
	Extend  Extend
	Lang    Lang
	Units   Units
}

// An Alert is an alert.
type Alert struct {
	Description string   `json:"description"`
	Expires     *Time    `json:"expires"`
	Regions     []string `json:"regions"`
	Severity    string   `json:"severity"`
	Time        *Time    `json:"time"`
	Title       string   `json:"title"`
	URI         string   `json:"uri"`
}

// A Currently is a current observation.
type Currently struct {
	ApparentTemperature  float64 `json:"apparentTemperature"`
	CloudCover           float64 `json:"cloudCover"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	Icon                 string  `json:"icon"`
	NearestStormBearing  float64 `json:"nearestStormBearing"`
	NearestStormDistance float64 `json:"nearestStormDistance"`
	Ozone                float64 `json:"ozone"`
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipProbability    float64 `json:"precipProbability"`
	Pressure             float64 `json:"pressure"`
	Summary              string  `json:"summary"`
	Temperature          float64 `json:"temperature"`
	Time                 *Time   `json:"time"`
	UVIndex              float64 `json:"uvIndex"`
	Visibility           float64 `json:"visibility"`
	WindBearing          float64 `json:"windBearing"`
	WindGust             float64 `json:"windGust"`
	WindSpeed            float64 `json:"windSpeed"`
}

// DailyData are daily forecast data.
type DailyData struct {
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime *Time   `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  *Time   `json:"apparentTemperatureLowTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  *Time   `json:"apparentTemperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  float64 `json:"apparentTemperatureMinTime"`
	CloudCover                  float64 `json:"cloudCover"`
	DewPoint                    float64 `json:"dewPoint"`
	Humidity                    float64 `json:"humidity"`
	Icon                        string  `json:"icon"`
	MoonPhase                   float64 `json:"moonPhase"`
	Ozone                       float64 `json:"ozone"`
	PrecipIntensity             float64 `json:"precipIntensity"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      *Time   `json:"precipIntensityMaxTime"`
	PrecipProbability           float64 `json:"precipProbability"`
	PrecipType                  string  `json:"precipType"`
	Pressure                    float64 `json:"pressure"`
	Summary                     string  `json:"summary"`
	SunriseTime                 *Time   `json:"sunriseTime"`
	SunsetTime                  *Time   `json:"sunsetTime"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         *Time   `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          *Time   `json:"temperatureLowTime"`
	TemperatureMax              float64 `json:"temperatureMax"`
	TemperatureMaxTime          *Time   `json:"temperatureMaxTime"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          *Time   `json:"temperatureMinTime"`
	Time                        *Time   `json:"time"`
	UVIndex                     float64 `json:"uvIndex"`
	UVIndexTime                 *Time   `json:"uvIndexTime"`
	Visibility                  float64 `json:"visibility"`
	WindBearing                 float64 `json:"windBearing"`
	WindGust                    float64 `json:"windGust"`
	WindGustTime                *Time   `json:"windGustTime"`
	WindSpeed                   float64 `json:"windSpeed"`
}

// A Daily is a daily forecast.
type Daily struct {
	Data    []*DailyData `json:"data"`
	Icon    string       `json:"icon"`
	Summary string       `json:"summary"`
}

// Flags are forecast flags.
type Flags struct {
	DarkSkyUnavailable interface{} `json:"darksky-unavailable"`
	NearestStation     float64     `json:"nearest-station"`
	Sources            []string    `json:"sources"`
	Units              Units       `json:"units"`
}

// HourlyData are hourly forecast data.
type HourlyData struct {
	ApparentTemperature float64 `json:"apparentTemperature"`
	CloudCover          float64 `json:"cloudCover"`
	DewPoint            float64 `json:"dewPoint"`
	Humidity            float64 `json:"humidity"`
	Icon                string  `json:"icon"`
	Ozone               float64 `json:"ozone"`
	PrecipIntensity     float64 `json:"precipIntensity"`
	PrecipProbability   float64 `json:"precipProbability"`
	PrecipType          string  `json:"precipType"`
	Pressure            float64 `json:"pressure"`
	Summary             string  `json:"summary"`
	Temperature         float64 `json:"temperature"`
	Time                *Time   `json:"time"`
	UVIndex             float64 `json:"uvIndex"`
	Visibility          float64 `json:"visibility"`
	WindBearing         float64 `json:"windBearing"`
	WindGust            float64 `json:"windGust"`
	WindSpeed           float64 `json:"windSpeed"`
}

// An Hourly is an hourly forecast.
type Hourly struct {
	Data    []*HourlyData `json:"data"`
	Icon    string        `json:"icon"`
	Summary string        `json:"summary"`
}

// MinutelyData are minutely forecast data.
type MinutelyData struct {
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipIntensityError float64 `json:"precipIntensityError"`
	PrecipProbability    float64 `json:"precipProbability"`
	PrecipType           string  `json:"precipType"`
	Time                 *Time   `json:"time"`
}

// A Minutely is a minutely forecast.
type Minutely struct {
	Data    []*MinutelyData `json:"data"`
	Icon    string          `json:"icon"`
	Summary string          `json:"summary"`
}

// A Forecast is a forecast.
type Forecast struct {
	Alerts    []*Alert   `json:"alerts"`
	Currently *Currently `json:"currently"`
	Daily     *Daily     `json:"daily"`
	Flags     *Flags     `json:"flags"`
	Hourly    *Hourly    `json:"hourly"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Minutely  *Minutely  `json:"minutely"`
	Offset    float64    `json:"offset"`
	Timezone  string     `json:"timezone"`
}

// Forecast returns the forecast for latitude and longitude at time t. If t is
// nil or zero then a forecast request is sent. If t is non-nil and non-zero
// then a time machine request is sent.
func (c *Client) Forecast(ctx context.Context, latitude, longitude float64, t *Time, options *ForecastOptions) (*Forecast, error) {
	urlStr := fmt.Sprintf("%s/forecast/%s/%f,%f", c.baseURL, c.key, latitude, longitude)
	if t != nil && !t.IsZero() {
		urlStr += "," + strconv.FormatInt(t.Unix(), 10)
	}

	if options != nil {
		values := url.Values{}
		if len(options.Exclude) != 0 {
			blockStrs := make([]string, len(options.Exclude))
			for i, block := range options.Exclude {
				blockStrs[i] = string(block)
			}
			sort.Strings(blockStrs)
			values.Set("exclude", strings.Join(blockStrs, ","))
		}
		if options.Extend != "" {
			values.Set("extend", string(options.Extend))
		}
		if options.Lang != "" {
			values.Set("lang", string(options.Lang))
		}
		if options.Units != "" {
			values.Set("units", string(options.Units))
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

	if c.responseMetadataCallback != nil {
		var forecastAPICalls int
		if facStr := resp.Header.Get("X-Forecast-API-Calls"); facStr != "" {
			if fac, err := strconv.ParseInt(facStr, 10, 64); err == nil {
				forecastAPICalls = int(fac)
			}
		}
		var responseTime time.Duration
		if rtStr := resp.Header.Get("X-Response-Time"); rtStr != "" {
			if rt, err := time.ParseDuration(rtStr); err == nil {
				responseTime = rt
			}
		}
		c.responseMetadataCallback(&ResponseMetadata{
			StatusCode:       resp.StatusCode,
			ForecastAPICalls: forecastAPICalls,
			ResponseTime:     responseTime,
		})
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		respBody, err := ioutil.ReadAll(resp.Body)
		e := &Error{
			Request:      req,
			Response:     resp,
			ResponseBody: respBody,
		}
		if err == nil {
			_ = json.Unmarshal(respBody, &e.Details)
		}
		return nil, e
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
