package api_getfood

import (
	"GoChiLeMaWails/src/database"
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
	"sort"
)

type RequestBody struct {
	Page     int  `json:"page"`
	PerPage  int  `json:"perPage"`
	Reversed bool `json:"reversed"`
}

func GetfoodRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get page from request body
	var body RequestBody
	err := route_utils.ReadJsonBody(r, &body)
	if err != nil {
		errorCode := route_utils.ApiErrorCode["FReadRequestBody"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	// Load foodtable
	foodTable := database.NewFoodTable("data/food.json")
	// If reversed
	if body.Reversed {
		sort.Slice(foodTable.Items, func(i, j int) bool {
			return foodTable.Items[i].Id > foodTable.Items[j].Id
		})
	}
	// Check if page is invalid
	itemsPerPage := body.PerPage
	totalPages := len(foodTable.Items) / itemsPerPage
	if len(foodTable.Items)%itemsPerPage != 0 {
		totalPages++
	}
	if body.Page < 0 || body.Page >= totalPages {
		errorCode := route_utils.ApiErrorCode["IPage"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get items
	startIndex := body.Page * itemsPerPage
	endIndex := startIndex + itemsPerPage
	if endIndex > len(foodTable.Items) {
		endIndex = len(foodTable.Items)
	}
	pageItems := foodTable.Items[startIndex:endIndex]
	// Return items
	food := make(map[string]interface{})
	food["items"] = pageItems
	food["totalPages"] = totalPages
	route_utils.WriteJSONResponse(w, food)
}
