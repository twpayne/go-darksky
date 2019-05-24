package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"text/template"

	darksky "github.com/twpayne/go-darksky"
)

var (
	latitude   = flag.Float64("latitude", 34.021944, "latitude")
	longitude  = flag.Float64("longitude", -118.481389, "longitude")
	timeVal    = flag.Int("time", 0, "time")
	exclude    = flag.String("exclude", "", "exclude")
	extend     = flag.String("extend", "", "extend")
	lang       = flag.String("lang", "", "language")
	units      = flag.String("units", "", "units")
	packageVal = flag.String("package", "dstest", "package")
)

var ouptutTemplate = template.Must(template.New("output").Parse(`// Automatically generated file. DO NOT EDIT.

package {{ .Package }}

import (
	"time"

	darksky "github.com/twpayne/go-darksky"
)

func init() {
	request := Request{
		Latitude:  {{ .Latitude }},
		Longitude: {{ .Longitude }},
		Time: darksky.Time{
			Time: {{ if eq .Time 0 }}time.Time{}{{ else }}time.Unix({{ .Time }}, 0){{ end }},
		},
		Exclude: "{{ .Exclude }}",
		Extend:  darksky.Extend("{{ .Extend }}"),
		Lang:    darksky.Lang("{{ .Lang }}"),
		Units:   darksky.Units("{{ .Units }}"),
	}
	// {{ .URL }}
	forecastStr := {{ .BackquotedJSON }}
	defaultForecasts[request] = forecastStr
}
`))

func run() error {
	flag.Parse()
	key := os.Getenv("DARKSKY_KEY")
	u, err := url.Parse(darksky.DefaultBaseURL)
	if err != nil {
		return err
	}
	u.Path += fmt.Sprintf("/forecast/%s/%f,%f", key, *latitude, *longitude)
	if *timeVal != 0 {
		u.Path += fmt.Sprintf(",%d", *timeVal)
	}
	q := u.Query()
	if *exclude != "" {
		blocks := strings.Split(*exclude, ",")
		sort.Strings(blocks)
		*exclude = strings.Join(blocks, ",")
		q.Set("exclude", *exclude)
	}
	if *extend != "" {
		q.Set("extend", *extend)
	}
	if *lang != "" {
		q.Set("lang", *lang)
	} else {
		*lang = string(darksky.DefaultLang)
	}
	if *units != "" {
		q.Set("units", *units)
	} else {
		*units = string(darksky.DefaultUnits)
	}
	if len(q) != 0 {
		u.RawQuery = q.Encode()
	}
	urlStr := u.String()
	//nolint: gosec
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return ouptutTemplate.Execute(os.Stdout, map[string]interface{}{
		"BackquotedJSON": fmt.Sprintf("`%s`", jsonData),
		"Exclude":        *exclude,
		"Extend":         *extend,
		"Lang":           *lang,
		"Latitude":       *latitude,
		"Longitude":      *longitude,
		"Package":        *packageVal,
		"Time":           *timeVal,
		"Units":          *units,
		"URL":            strings.Replace(urlStr, key, "${DARKSKY_KEY}", -1),
	})
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
