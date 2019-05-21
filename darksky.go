// Package darksky implements a client for the Dark Sky weather forecasting API.
// See https://darksky.net/dev.
package darksky

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// Attribution is the attribution that must be displayed.
	Attribution = "Powered by Dark Sky"

	// AttributionURL is the URL that the displayed attribution must link to.
	AttributionURL = "https://darksky.net/poweredby/"

	// DefaultBaseURL is the default base URL.
	DefaultBaseURL = "https://api.darksky.net"

	// DefaultLang is the default lang.
	DefaultLang = LangEN

	// DefaultUnits are the default units.
	DefaultUnits = UnitsUS
)

// An Error is an error.
type Error struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
	Details      struct {
		Code     int    `json:"code"`
		ErrorStr string `json:"error"`
	}
}

// ResponseMetadata are extra metadata associated with a response.
type ResponseMetadata struct {
	StatusCode       int
	ForecastAPICalls int
	ResponseTime     time.Duration
}

// A ResponseMetadataCallback is function that receives a ResponseMetadata.
type ResponseMetadataCallback func(*ResponseMetadata)

// A Client is a Dark Sky Client.
type Client struct {
	httpClient               *http.Client
	baseURL                  string
	key                      string
	responseMetadataCallback ResponseMetadataCallback
}

// A ClientOption sets an option on a Client.
type ClientOption func(*Client)

// WithHTTPClient sets the HTTP Client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets the base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithKey sets the key.
func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.key = key
	}
}

// WithResponseMetadataCallback sets the response metadata callback.
func WithResponseMetadataCallback(rmc ResponseMetadataCallback) ClientOption {
	return func(c *Client) {
		c.responseMetadataCallback = rmc
	}
}

// NewClient returns a new Client.
func NewClient(options ...ClientOption) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    DefaultBaseURL,
	}
	for _, o := range options {
		o(c)
	}
	return c
}

func (e *Error) Error() string {
	if e.Details.ErrorStr != "" {
		return e.Details.ErrorStr
	}
	s := fmt.Sprintf("%s: %d %s", e.Request.URL, e.Response.StatusCode, http.StatusText(e.Response.StatusCode))
	if len(e.ResponseBody) != 0 {
		s += ": " + string(e.ResponseBody)
	}
	return s
}
