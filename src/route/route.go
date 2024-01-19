package route

import (
	"GoChiLeMaWails/src/global"
	api_contactus "GoChiLeMaWails/src/route/api/contactus"
	api_csort "GoChiLeMaWails/src/route/api/csort"
	api_deletefood "GoChiLeMaWails/src/route/api/deletefood"
	api_getfood "GoChiLeMaWails/src/route/api/getfood"
	api_gosort "GoChiLeMaWails/src/route/api/gosort"
	api_ip "GoChiLeMaWails/src/route/api/ip"
	api_osinfo "GoChiLeMaWails/src/route/api/osinfo"
	api_recommendfood "GoChiLeMaWails/src/route/api/recommendfood"
	api_recordfood "GoChiLeMaWails/src/route/api/recordfood"
	api_weather "GoChiLeMaWails/src/route/api/weather"
	"fmt"
	"net/http"
)

var RouteHandlers map[string]http.Handler

func Init() {
	RouteHandlers = make(map[string]http.Handler)
	// Register routes here
	RegisterRouteHandlerFunc("/api/contactus", api_contactus.ContactusRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/csort", api_csort.CsortRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/deletefood", api_deletefood.DeletefoodRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/getfood", api_getfood.GetfoodRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/gosort", api_gosort.GosortRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/ip", api_ip.IpRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/osinfo", api_osinfo.OsinfoRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/recommendfood", api_recommendfood.RecommendfoodRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/recordfood", api_recordfood.RecordfoodRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/weather", api_weather.WeatherRouteHandlerFunc)
	fmt.Println("Registered routes:", RouteHandlers)
	// Handle all routes and start the server
	go HandleAllRoutes()
	go http.ListenAndServe(fmt.Sprintf(":%d", global.API_PORT), nil)
	fmt.Println("Listening on port", global.API_PORT)
}

// RegisterRouteHandler registers a route handler for a given route.
func RegisterRouteHandler(route string, handler http.Handler) {
	RouteHandlers[route] = handler
}

// RegisterRouteHandlerFunc registers a route handler function for a given route.
func RegisterRouteHandlerFunc(route string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	handler := http.HandlerFunc(handlerFunc)
	RegisterRouteHandler(route, handler)
}

// HandleAllRoutes handles all routes registered with RegisterRouteHandler or RegisterRouteHandlerFunc.
func HandleAllRoutes() {
	for route, handler := range RouteHandlers {
		fmt.Println("Handling route:", route)
		go http.Handle(route, handler)
	}
}
