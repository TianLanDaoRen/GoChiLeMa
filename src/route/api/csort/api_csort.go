package api_csort

/*
#cgo LDFLAGS: -static
#include "csort.h"
*/
import "C"

import (
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
	"unsafe"
)

type RequestBody struct {
	Numbers []int32 `json:"numbers"`
}

func CsortRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	// Get numbers from request body
	var body RequestBody
	err := route_utils.ReadJsonBody(r, &body)
	if err != nil {
		errorCode := route_utils.ApiErrorCode["FReadRequestBody"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	// Call C function
	lNumbers := len(body.Numbers)
	cNumbers := (*C.int)(unsafe.Pointer(&body.Numbers[0]))
	C.Sort(cNumbers, C.int(lNumbers))
	sorted := (*[1 << 30]C.int)(unsafe.Pointer(cNumbers))
	sortedSlice := sorted[:lNumbers]
	sort := make(map[string]interface{})
	sort["numbers"] = sortedSlice
	route_utils.WriteJSONResponse(w, sort)
}
