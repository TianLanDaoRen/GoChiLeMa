package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func GetExternalIPv4() (string, error) {
	resp, err := http.Get("http://api.ipify.org")
	if err != nil {
		fmt.Println("Failed to get external IPv4 address:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", err
	}

	return string(body), nil
}

func GetLocationInfo(ip string) (location Location, err error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get cityname:", err)
		return Location{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return Location{}, err
	}

	// Unmarshal json
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println("Failed to unmarshal json:", err)
		return Location{}, err
	}

	return location, nil
}
