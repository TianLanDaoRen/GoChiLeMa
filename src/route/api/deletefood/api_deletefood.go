package api_deletefood

import (
	"GoChiLeMaWails/src/database"
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
)

type RequestBody struct {
	Ids []int `json:"ids"`
}

func DeletefoodRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get ids from request body
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
	// Delete items
	foodTable.RemoveItemsByIds(body.Ids)
	// Save foodtable
	foodTable.Save()
	// Return items
	food := make(map[string]interface{})
	food["message"] = "删除数据成功！"
	route_utils.WriteJSONResponse(w, food)
}
