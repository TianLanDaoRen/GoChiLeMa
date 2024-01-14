package api_contactus

import (
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
)

func ContactusRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		invalid := route_utils.MakeDefaultJSONResponse(2401, "Invalid request")
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	contact := make(map[string]interface{})
	contact["email"] = "moriartylimitter@outlook.com"
	contact["vx"] = "TianLanDaoRen"
	route_utils.WriteJSONResponse(w, contact)
}
