package api_ip

import (
	"GoChiLeMaWails/src/global"
	"GoChiLeMaWails/src/location"
	route_utils "GoChiLeMaWails/src/route/utils"
	"fmt"
	"net/http"
)

func IpRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		invalid := route_utils.MakeDefaultJSONResponse(2401, "Invalid request")
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get IP
	if global.IP == "" {
		iip, err := location.GetExternalIPv4()
		if err != nil {
			failed := route_utils.MakeDefaultJSONResponse(2402, "Failed to get external IPv4")
			route_utils.WriteJSONResponse(w, failed)
			return
		}
		global.IP = iip
	}
	// Get location info
	if global.Lat == 0 || global.Lon == 0 || global.City == "" || global.Timezone == "" || global.ISP == "" || global.AS == "" {
		loc, err := location.GetLocationInfo(global.IP)
		if err != nil {
			failed := route_utils.MakeDefaultJSONResponse(2403, "Failed to get location information")
			route_utils.WriteJSONResponse(w, failed)
			return
		}
		global.Lat = loc.Lat
		global.Lon = loc.Lon
		global.City = fmt.Sprintf("%s %s %s", loc.Country, loc.RegionName, loc.City)
		global.Timezone = loc.Timezone
		global.ISP = loc.Isp
		global.AS = loc.As
	}
	ip := make(map[string]interface{})
	ip["ip"] = global.IP
	ip["city"] = global.City
	ip["lat"] = global.Lat
	ip["lon"] = global.Lon
	ip["timezone"] = global.Timezone
	ip["isp"] = global.ISP
	ip["as"] = global.AS
	route_utils.WriteJSONResponse(w, ip)
}
