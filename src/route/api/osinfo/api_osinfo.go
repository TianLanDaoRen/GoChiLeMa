package api_osinfo

/*
#cgo LDFLAGS: -static
#include "osinfo.h"
*/
import "C"

import (
	"GoChiLeMaWails/src/route/utils"
	"net/http"
)

func OsinfoRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// make default json response
	os := make(map[string]interface{})
	os["info"] = "undefined"

	// Check valid call
	valid := utils.CheckValidRequest(r)
	if !valid {
		utils.WriteJSONResponse(w, os)
		return
	}
	// Call C function
	osinfo := C.GetOsVersion()
	os["info"] = C.GoString(osinfo)
	utils.WriteJSONResponse(w, os)
}
