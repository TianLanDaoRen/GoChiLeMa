package route

import (
	"GoChiLeMa/src/config"
	api_contactus "GoChiLeMa/src/route/api/contactus"
	api_ip "GoChiLeMa/src/route/api/ip"
	api_weather "GoChiLeMa/src/route/api/weather"
	"fmt"
	"net/http"
)

var RouteHandlers map[string]http.Handler

func Init() {
	RouteHandlers = make(map[string]http.Handler)
	// Register routes here
	RegisterRouteHandlerFunc("/api/contactus", api_contactus.ContactusRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/ip", api_ip.IpRouteHandlerFunc)
	RegisterRouteHandlerFunc("/api/weather", api_weather.WeatherRouteHandlerFunc)
	fmt.Println("Registered routes:", RouteHandlers)
	// Handle all routes and start the server
	go HandleAllRoutes()
	go http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), nil)
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
