package api_gosort

import (
	route_utils "GoChiLeMaWails/src/route/utils"
	"net/http"
	"sort"
	"time"
)

type RequestBody struct {
	Sign    string  `json:"sign"`
	Sn      float64 `json:"sn"`
	Ts      float64 `json:"ts"`
	Numbers []int32 `json:"numbers"`
}

func Sort(numbers []int32) {
	// Record start time
	startTime := time.Now()
	// Use golang sort.Slice
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	// Record end time
	endTime := time.Now()
	// Calculate time elapsed
	timeElapsed := endTime.Sub(startTime)
	// Print time elapsed
	println("Go sort Time elapsed:", timeElapsed.Milliseconds())
}

func GosortRouteHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Check valid call
	valid := route_utils.CheckValidRequest(r)
	if !valid {
		invalid := route_utils.MakeDefaultJSONResponse(2401, "Invalid request")
		route_utils.WriteJSONResponse(w, invalid)
		return
	}
	// Get numbers from request body
	var body RequestBody
	err := route_utils.ReadJsonBody(r, &body)
	if err != nil {
		failed := route_utils.MakeDefaultJSONResponse(2400, "Failed to read request body")
		route_utils.WriteJSONResponse(w, failed)
		return
	}
	// Call go function
	Sort(body.Numbers)
	sort := make(map[string]interface{})
	sort["numbers"] = body.Numbers
	route_utils.WriteJSONResponse(w, sort)
}
