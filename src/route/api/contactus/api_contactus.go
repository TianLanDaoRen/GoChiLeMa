package api_contactus

import (
	"GoChiLeMaWails/src/route/utils"
	"net/http"
)

func ContactusRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// make default json response
	contact := make(map[string]interface{})
	contact["email"] = "undefined"
	contact["vx"] = "undefined"

	// Check valid call
	valid := utils.CheckValidRequest(r)
	if !valid {
		utils.WriteJSONResponse(w, contact)
		return
	}
	contact["email"] = "moriartylimitter@outlook.com"
	contact["vx"] = "TianLanDaoRen"

	utils.WriteJSONResponse(w, contact)
}
