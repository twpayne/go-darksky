// Package dstest implements a mock Dark Sky server for testing.
package dstest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	darksky "github.com/twpayne/go-darksky"
)

// DefaultKey is the default key.
const DefaultKey = "key"

var (
	errOutOfRange = errors.New("out of range")

	//nolint: gochecknoglobals
	defaultForecasts = make(map[Request]string)
)

// A Request contains parameters for a request.
type Request struct {
	Latitude  float64
	Longitude float64
	Time      darksky.Time
	Exclude   string
	Extend    darksky.Extend
	Lang      darksky.Lang
	Units     darksky.Units
}

// A Server is a mock server.
type Server struct {
	*httptest.Server
	chi.Router
	Key       string
	Forecasts map[Request]string
}

// An Option sets an option on a Server.
type Option func(*Server)

// WithDefaultForecasts returns an option that adds all default forecasts to a
// Server.
func WithDefaultForecasts() Option {
	return func(s *Server) {
		for request, forecast := range defaultForecasts {
			s.Forecasts[request] = forecast
		}
	}
}

// WithForecast returns an option that adds a forecastStr as a response to
// request on a Server.
func WithForecast(request Request, forecastStr string) Option {
	return func(s *Server) {
		s.Forecasts[request] = forecastStr
	}
}

// NewServer returns a new Server.
func NewServer(options ...Option) *Server {
	router := chi.NewRouter()
	s := &Server{
		Server:    httptest.NewServer(router),
		Router:    router,
		Key:       DefaultKey,
		Forecasts: make(map[Request]string),
	}
	router.Get("/forecast/{key}/{latitude},{longitude}", s.handleForecast)
	router.Get("/forecast/{key}/{latitude},{longitude},{time}", s.handleForecast)
	for _, o := range options {
		o(s)
	}
	return s
}

// NewClient returns a new darksky.Client that connects to s.
func (s *Server) NewClient(options ...darksky.ClientOption) *darksky.Client {
	return darksky.NewClient(
		append([]darksky.ClientOption{
			darksky.WithBaseURL(s.URL),
			darksky.WithHTTPClient(s.Client()),
			darksky.WithKey(s.Key),
		}, options...)...,
	)
}

func (s *Server) handleForecast(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "key") != s.Key {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	request := Request{
		Lang:  darksky.LangEN,
		Units: darksky.UnitsUS,
	}

	var err error
	request.Latitude, err = parseFloatFromURLParam(r, "latitude", -90, 90)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Longitude, err = parseFloatFromURLParam(r, "longitude", -180, 180)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if timeStr := chi.URLParam(r, "time"); timeStr != "" {
		timeVal, err := strconv.Atoi(timeStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		request.Time.Time = time.Unix(int64(timeVal), 0)
	}

	query := r.URL.Query()
	if exclude := query.Get("exclude"); exclude != "" {
		request.Exclude = exclude
	}
	if extend := query.Get("extend"); extend != "" {
		request.Extend = darksky.Extend(extend)
	}
	if lang := query.Get("lang"); lang != "" {
		request.Lang = darksky.Lang(lang)
	}
	if units := query.Get("units"); units != "" {
		request.Units = darksky.Units(units)
	}

	forecast, ok := s.Forecasts[request]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, _ = w.Write([]byte(forecast))
}

// Options returns a new darksky.ForecastOptions constructed from r.
func (r *Request) Options() *darksky.ForecastOptions {
	if r.Exclude == "" && r.Extend == darksky.DefaultExtend && r.Lang == darksky.DefaultLang && r.Units == darksky.DefaultUnits {
		return nil
	}
	o := &darksky.ForecastOptions{}
	if r.Exclude != "" {
		for _, block := range strings.Split(r.Exclude, ",") {
			o.Exclude = append(o.Exclude, darksky.Block(block))
		}
	}
	if r.Extend != darksky.DefaultExtend {
		o.Extend = r.Extend
	}
	if r.Lang != darksky.DefaultLang {
		o.Lang = r.Lang
	}
	if r.Units != darksky.DefaultUnits {
		o.Units = r.Units
	}
	return o
}

func parseFloatFromURLParam(r *http.Request, key string, min, max float64) (float64, error) {
	x, err := strconv.ParseFloat(chi.URLParam(r, key), 64)
	if err != nil {
		return 0, err
	}
	if x < min || max < x {
		return 0, errOutOfRange
	}
	return x, nil
}
