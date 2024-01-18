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
		errorCode := route_utils.ApiErrorCode["IRequest"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	// Get IP
	if global.IP == "" {
		iip, err := location.GetExternalIPv4()
		if err != nil {
			errorCode := route_utils.ApiErrorCode["FGetExternalIpv4"]
			apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
			route_utils.WriteJSONResponse(w, apiError)
			return
		}
		global.IP = iip
	}
	// Get location info
	if global.Lat == 0 || global.Lon == 0 || global.City == "" || global.Timezone == "" || global.ISP == "" || global.AS == "" {
		loc, err := location.GetLocationInfo(global.IP)
		if err != nil {
			errorCode := route_utils.ApiErrorCode["FGetLocationInfo"]
			apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
			route_utils.WriteJSONResponse(w, apiError)
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
