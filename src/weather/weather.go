package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset float64 `json:"timezone_offset"`
	Current        Current `json:"current"`
}

type Current struct {
	Dt         float64       `json:"dt"`
	Sunrise    float64       `json:"sunrise"`
	Sunset     float64       `json:"sunset"`
	Temp       float64       `json:"temp"`
	FeelsLike  float64       `json:"feels_like"`
	Pressure   float64       `json:"pressure"`
	Humidity   float64       `json:"humidity"`
	DewPoint   float64       `json:"dew_point"`
	Uvi        float64       `json:"uvi"`
	Clouds     float64       `json:"clouds"`
	Visibility float64       `json:"visibility"`
	WindSpeed  float64       `json:"wind_speed"`
	WindDeg    float64       `json:"wind_deg"`
	WindGust   float64       `json:"wind_gust"`
	Weather    []WeatherInfo `json:"weather"`
}

type WeatherInfo struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherApi struct {
	Lat, Lon, Exclude, ApiKey string
}

func (api *WeatherApi) GetWeather() (Weather, bool) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&exclude=%s&appid=%s&units=metric&lang=zh_cn", api.Lat, api.Lon, api.Exclude, api.ApiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return Weather{}, false
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return Weather{}, false
	}

	// Unmarshal the JSON bytes into a value of the appropriate type.
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Weather{}, false
	}

	return weather, true
}
