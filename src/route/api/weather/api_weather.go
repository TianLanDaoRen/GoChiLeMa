package api_weather

import (
	"GoChiLeMaWails/src/global"
	"GoChiLeMaWails/src/location"
	route_utils "GoChiLeMaWails/src/route/utils"
	WEATHER "GoChiLeMaWails/src/weather"
	"fmt"
	"net/http"
	"time"
)

func WeatherRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		invalid := route_utils.MakeDefaultJSONResponse(2401, "Invalid request")
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get IP
	if global.IP == "" {
		ip, err := location.GetExternalIPv4()
		if err != nil {
			failed := route_utils.MakeDefaultJSONResponse(2402, "Failed to get external IPv4")
			route_utils.WriteJSONResponse(w, failed)
			return
		}
		global.IP = ip
	}
	// Get location info
	if global.Lat == 0 || global.Lon == 0 {
		loc, err := location.GetLocationInfo(global.IP)
		if err != nil {
			failed := route_utils.MakeDefaultJSONResponse(2403, "Failed to get location information")
			route_utils.WriteJSONResponse(w, failed)
			return
		}
		global.Lat = loc.Lat
		global.Lon = loc.Lon
	}
	// Get weather info
	if global.WeatherInfo == nil {
		weatherApi := WEATHER.WeatherApi{
			Lat:     fmt.Sprintf("%f", global.Lat),
			Lon:     fmt.Sprintf("%f", global.Lon),
			Exclude: "minutely,hourly,daily,alerts",
			ApiKey:  "9dc2561b8fe88e3cb9697cfcd4bd770d",
		}
		weatherInfo, ok := weatherApi.GetWeather()
		if !ok {
			failed := route_utils.MakeDefaultJSONResponse(2404, "Failed to get weather information")
			route_utils.WriteJSONResponse(w, failed)
			return
		}
		global.WeatherInfo = &WEATHER.Weather{}
		*global.WeatherInfo = weatherInfo
	}
	timezoneOffset := int(global.WeatherInfo.TimezoneOffset / 3600)
	temp := fmt.Sprintf("%.2f", global.WeatherInfo.Current.Temp)
	desc := global.WeatherInfo.Current.Weather[0].Description
	humidity := int(global.WeatherInfo.Current.Humidity)
	uvi := global.WeatherInfo.Current.Uvi
	clouds := int(global.WeatherInfo.Current.Clouds)
	windSpeed := global.WeatherInfo.Current.WindSpeed
	sunriseTimestamp := int64(global.WeatherInfo.Current.Sunrise)
	sunrise := time.Unix(sunriseTimestamp, 0).UTC().Add(time.Hour * time.Duration(timezoneOffset)).Format("15:04:05")
	sunsetTimestamp := int64(global.WeatherInfo.Current.Sunset)
	sunset := time.Unix(sunsetTimestamp, 0).UTC().Add(time.Hour * time.Duration(timezoneOffset)).Format("15:04:05")
	description := fmt.Sprintf("%s，湿度%d%%，紫外线指数%.2f，云度%d%%，风速%.2fm/s，日出时间%s，日落时间%s", desc, humidity, uvi, clouds, windSpeed, sunrise, sunset)

	weather := make(map[string]interface{})
	weather["temp"] = temp
	weather["description"] = description
	route_utils.WriteJSONResponse(w, weather)
}
