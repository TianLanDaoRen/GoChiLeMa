package api_contactus

import (
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
)

func ContactusRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	contact := make(map[string]interface{})
	contact["email"] = "moriartylimitter@outlook.com"
	contact["vx"] = "TianLanDaoRen"
	route_utils.WriteJSONResponse(w, contact)
}
