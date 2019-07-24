// Automatically generated file. DO NOT EDIT.

package dstest

import (
	"time"

	darksky "github.com/trende-jp/go-darksky"
)

func init() {
	request := Request{
		Latitude:  34.0219,
		Longitude: -118.4814,
		Time: darksky.Time{
			Time: time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		Exclude: "",
		Extend:  darksky.Extend(""),
		Lang:    darksky.Lang("en"),
		Units:   darksky.Units("us"),
	}
	// https://api.darksky.net/forecast/${DARKSKY_KEY}/34.021900,-118.481400,1556668800
	forecastStr := `{"latitude":34.0219,"longitude":-118.4814,"timezone":"America/Los_Angeles","currently":{"time":1556668800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":62.56,"apparentTemperature":62.56,"dewPoint":48.82,"humidity":0.61,"pressure":1014.72,"windSpeed":4.79,"windGust":9.8,"windBearing":222,"cloudCover":0.82,"uvIndex":2,"visibility":8.84,"ozone":338.18},"hourly":{"summary":"Mostly cloudy until night.","icon":"partly-cloudy-day","data":[{"time":1556607600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":55.86,"apparentTemperature":55.86,"dewPoint":51.95,"humidity":0.87,"pressure":1012.91,"windSpeed":2.81,"windGust":3.11,"windBearing":293,"cloudCover":0.3,"uvIndex":0,"visibility":10,"ozone":345.02},{"time":1556611200,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":54.8,"apparentTemperature":54.8,"dewPoint":51.83,"humidity":0.9,"pressure":1013.07,"windSpeed":3.07,"windGust":4.06,"windBearing":283,"cloudCover":0.66,"uvIndex":0,"visibility":10,"ozone":345.03},{"time":1556614800,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":54.42,"apparentTemperature":54.42,"dewPoint":51.75,"humidity":0.91,"pressure":1013,"windSpeed":2.42,"windGust":4.67,"windBearing":276,"cloudCover":0.88,"uvIndex":0,"visibility":10,"ozone":345.64},{"time":1556618400,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":54.3,"apparentTemperature":54.3,"dewPoint":51.31,"humidity":0.9,"pressure":1013,"windSpeed":2.26,"windGust":3.66,"windBearing":292,"cloudCover":0.95,"uvIndex":0,"visibility":10,"ozone":347.37},{"time":1556622000,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":54.32,"apparentTemperature":54.32,"dewPoint":51.6,"humidity":0.9,"pressure":1013.15,"windSpeed":1.56,"windGust":2.96,"windBearing":309,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":349.61},{"time":1556625600,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":54.71,"apparentTemperature":54.71,"dewPoint":51.92,"humidity":0.9,"pressure":1013.33,"windSpeed":1.1,"windGust":2.55,"windBearing":307,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":351.55},{"time":1556629200,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0013,"precipProbability":0.1,"precipType":"rain","temperature":54.88,"apparentTemperature":54.88,"dewPoint":52.35,"humidity":0.91,"pressure":1013.59,"windSpeed":0.98,"windGust":1.87,"windBearing":250,"cloudCover":1,"uvIndex":0,"visibility":10,"ozone":353.15},{"time":1556632800,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":55.52,"apparentTemperature":55.52,"dewPoint":52.75,"humidity":0.9,"pressure":1014.32,"windSpeed":1.96,"windGust":2.77,"windBearing":249,"cloudCover":1,"uvIndex":0,"visibility":9.51,"ozone":354.42},{"time":1556636400,"summary":"Overcast","icon":"cloudy","precipIntensity":0.0003,"precipProbability":0.03,"precipType":"rain","temperature":56.41,"apparentTemperature":56.41,"dewPoint":52.86,"humidity":0.88,"pressure":1015.1,"windSpeed":1.88,"windGust":3.93,"windBearing":239,"cloudCover":1,"uvIndex":1,"visibility":10,"ozone":355},{"time":1556640000,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0042,"precipProbability":0.06,"precipType":"rain","temperature":57.61,"apparentTemperature":57.61,"dewPoint":52.44,"humidity":0.83,"pressure":1015.41,"windSpeed":4.6,"windGust":6.52,"windBearing":228,"cloudCover":0.65,"uvIndex":2,"visibility":10,"ozone":354.71},{"time":1556643600,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":58.76,"apparentTemperature":58.76,"dewPoint":52.37,"humidity":0.79,"pressure":1015.91,"windSpeed":3.16,"windGust":6.43,"windBearing":239,"cloudCover":1,"uvIndex":3,"visibility":9.68,"ozone":353.73},{"time":1556647200,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":59.43,"apparentTemperature":59.43,"dewPoint":51.99,"humidity":0.76,"pressure":1016.31,"windSpeed":3.6,"windGust":7.94,"windBearing":228,"cloudCover":1,"uvIndex":3,"visibility":10,"ozone":351.86},{"time":1556650800,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":60.59,"apparentTemperature":60.59,"dewPoint":51.8,"humidity":0.73,"pressure":1016.6,"windSpeed":4.19,"windGust":9.12,"windBearing":235,"cloudCover":1,"uvIndex":4,"visibility":9.7,"ozone":350.11},{"time":1556654400,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":61.47,"apparentTemperature":61.47,"dewPoint":51.28,"humidity":0.69,"pressure":1016.27,"windSpeed":4.4,"windGust":9.97,"windBearing":239,"cloudCover":0.99,"uvIndex":4,"visibility":9.43,"ozone":348.06},{"time":1556658000,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":62.12,"apparentTemperature":62.12,"dewPoint":50.78,"humidity":0.66,"pressure":1015.68,"windSpeed":4.92,"windGust":11.02,"windBearing":230,"cloudCover":0.8,"uvIndex":5,"visibility":9.33,"ozone":346.18},{"time":1556661600,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":62.55,"apparentTemperature":62.55,"dewPoint":50.06,"humidity":0.64,"pressure":1015.26,"windSpeed":4.32,"windGust":10.94,"windBearing":222,"cloudCover":0.72,"uvIndex":4,"visibility":9.25,"ozone":344.36},{"time":1556665200,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":62.52,"apparentTemperature":62.52,"dewPoint":49.5,"humidity":0.62,"pressure":1014.88,"windSpeed":4.93,"windGust":9.73,"windBearing":217,"cloudCover":0.89,"uvIndex":3,"visibility":8.51,"ozone":342.62},{"time":1556668800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":62.56,"apparentTemperature":62.56,"dewPoint":48.82,"humidity":0.61,"pressure":1014.72,"windSpeed":4.79,"windGust":9.8,"windBearing":222,"cloudCover":0.82,"uvIndex":2,"visibility":8.84,"ozone":338.18},{"time":1556672400,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":61.49,"apparentTemperature":61.49,"dewPoint":48.53,"humidity":0.62,"pressure":1014.66,"windSpeed":5.5,"windGust":8.91,"windBearing":250,"cloudCover":0.6,"uvIndex":1,"visibility":8.97,"ozone":337.74},{"time":1556676000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":60.37,"apparentTemperature":60.37,"dewPoint":47.92,"humidity":0.63,"pressure":1014.65,"windSpeed":5.09,"windGust":8.16,"windBearing":246,"cloudCover":0.47,"uvIndex":0,"visibility":9.53,"ozone":338.29},{"time":1556679600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":58.56,"apparentTemperature":58.56,"dewPoint":48.08,"humidity":0.68,"pressure":1014.95,"windSpeed":4.17,"windGust":6.05,"windBearing":252,"cloudCover":0.45,"uvIndex":0,"visibility":9.64,"ozone":339.25},{"time":1556683200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":57.5,"apparentTemperature":57.5,"dewPoint":47.92,"humidity":0.7,"pressure":1015.42,"windSpeed":2.94,"windGust":5.06,"windBearing":215,"cloudCover":0.28,"uvIndex":0,"visibility":9.8,"ozone":340.45},{"time":1556686800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":56.48,"apparentTemperature":56.48,"dewPoint":48.11,"humidity":0.74,"pressure":1015.67,"windSpeed":2.24,"windGust":3.93,"windBearing":304,"cloudCover":0.34,"uvIndex":0,"visibility":9.75,"ozone":342.08},{"time":1556690400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":55.4,"apparentTemperature":55.4,"dewPoint":47.97,"humidity":0.76,"pressure":1015.67,"windSpeed":1.7,"windGust":3.69,"windBearing":283,"cloudCover":0.23,"uvIndex":0,"visibility":10,"ozone":347.31}]},"daily":{"data":[{"time":1556607600,"summary":"Mostly cloudy throughout the day.","icon":"partly-cloudy-day","sunriseTime":1556629609,"sunsetTime":1556678270,"moonPhase":0.87,"precipIntensity":0.0003,"precipIntensityMax":0.0042,"precipIntensityMaxTime":1556640000,"precipProbability":0.18,"precipType":"rain","temperatureHigh":62.56,"temperatureHighTime":1556668800,"temperatureLow":51.58,"temperatureLowTime":1556719200,"apparentTemperatureHigh":62.56,"apparentTemperatureHighTime":1556668800,"apparentTemperatureLow":51.58,"apparentTemperatureLowTime":1556719200,"dewPoint":50.66,"humidity":0.77,"pressure":1014.73,"windSpeed":2.94,"windGust":11.02,"windGustTime":1556658000,"windBearing":246,"cloudCover":0.75,"uvIndex":5,"uvIndexTime":1556658000,"visibility":9.66,"ozone":346.74,"temperatureMin":54.3,"temperatureMinTime":1556618400,"temperatureMax":62.56,"temperatureMaxTime":1556668800,"apparentTemperatureMin":54.3,"apparentTemperatureMinTime":1556618400,"apparentTemperatureMax":62.56,"apparentTemperatureMaxTime":1556668800}]},"flags":{"sources":["cmc","gfs","hrrr","icon","isd","madis","nam","sref"],"nearest-station":1.42,"units":"us"},"offset":-7}`
	defaultForecasts[request] = forecastStr
}
