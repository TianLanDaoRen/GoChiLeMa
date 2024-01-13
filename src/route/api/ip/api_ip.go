package api_ip

import (
	"GoChiLeMaWails/src/global"
	"GoChiLeMaWails/src/location"
	"GoChiLeMaWails/src/route/utils"
	"fmt"
	"net/http"
)

func IpRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// make default json response
	ip := make(map[string]interface{})
	ip["ip"] = "undefined"
	ip["city"] = "undefined"
	ip["lat"] = "undefined"
	ip["lon"] = "undefined"
	ip["timezone"] = "undefined"
	ip["isp"] = "undefined"
	ip["as"] = "undefined"

	// Check valid call
	valid := utils.CheckValidRequest(r)
	if !valid {
		utils.WriteJSONResponse(w, ip)
		return
	}
	// Get IP
	if global.IP == "" {
		iip, err := location.GetExternalIPv4()
		if err != nil {
			fmt.Println("Failed to get external IPv4")
			utils.WriteJSONResponse(w, ip)
			return
		}
		global.IP = iip
	}
	// Get location info
	if global.Lat == 0 || global.Lon == 0 || global.City == "" || global.Timezone == "" || global.ISP == "" || global.AS == "" {
		loc, err := location.GetLocationInfo(global.IP)
		if err != nil {
			fmt.Println("Failed to get location information")
			utils.WriteJSONResponse(w, ip)
			return
		}
		global.Lat = loc.Lat
		global.Lon = loc.Lon
		global.City = fmt.Sprintf("%s %s %s", loc.Country, loc.RegionName, loc.City)
		global.Timezone = loc.Timezone
		global.ISP = loc.Isp
		global.AS = loc.As
	}
	ip["ip"] = global.IP
	ip["city"] = global.City
	ip["lat"] = global.Lat
	ip["lon"] = global.Lon
	ip["timezone"] = global.Timezone
	ip["isp"] = global.ISP
	ip["as"] = global.AS
	utils.WriteJSONResponse(w, ip)
}
