package api_osinfo

/*
#cgo LDFLAGS: -static
#include "osinfo.h"
*/
import "C"

import (
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
)

func OsinfoRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Call C function
	os := make(map[string]interface{})
	osinfo := C.GetOsVersion()
	os["info"] = C.GoString(osinfo)
	route_utils.WriteJSONResponse(w, os)
}
