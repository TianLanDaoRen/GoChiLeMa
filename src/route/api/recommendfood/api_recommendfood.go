package api_recommendfood

import (
	"GoChiLeMaWails/src/database"
	"GoChiLeMaWails/src/markov"
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
)

type RequestBody struct {
	N int `json:"n"`
}

func RecommendfoodRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
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
	// Check if foodtable has item
	if len(foodTable.Items) <= 0 {
		errorCode := route_utils.ApiErrorCode["IRequest"]
		apiError := route_utils.MakeDefaultJSONResponse(errorCode.Code, errorCode.Msg)
		route_utils.WriteJSONResponse(w, apiError)
		return
	}
	// Run Markov prediction model
	foodList := make([]string, 0, len(foodTable.Items))
	for _, item := range foodTable.Items {
		foodList = append(foodList, item.Food)
	}
	markovModel := markov.NewMarkovModel(foodList)
	next := markovModel.GetNextPiNEvent(body.N)
	// Return items
	food := make(map[string]interface{})
	food["recommend"] = next
	route_utils.WriteJSONResponse(w, food)
}
