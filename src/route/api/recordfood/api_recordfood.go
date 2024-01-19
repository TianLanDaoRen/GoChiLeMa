package api_recordfood

import (
	"GoChiLeMaWails/src/database"
	route_utils "GoChiLeMaWails/src/route/utils"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
	Food  string `json:"food"`
}

func RecordfoodRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
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
	// Check if date is empty
	if body.Year == "" || body.Month == "" || body.Day == "" {
		errorCode := route_utils.ApiErrorCode["IDate"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Check if food is empty
	if body.Food == "" {
		errorCode := route_utils.ApiErrorCode["IFood"]
		invalid := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Append item
	foodTable := database.NewFoodTable("data/food.json")
	foodItem := database.FoodItem{
		Id:   foodTable.GetNextId(),
		Date: fmt.Sprintf("%s-%s-%s", body.Year, body.Month, body.Day),
		Food: body.Food,
	}
	foodTable.AppendItem(foodItem)
	// Save foodtable
	foodTable.Save()

	record := make(map[string]interface{})
	record["message"] = "记录数据成功！"
	route_utils.WriteJSONResponse(w, record)
}
