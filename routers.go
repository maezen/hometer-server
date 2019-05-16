/*
 * Hometer
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"fmt"
	"hometer-server/api"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetCurrentHumidity",
		strings.ToUpper("Get"),
		"/v1/humditity",
		api.GetCurrentHumidity,
	},

	Route{
		"GetLastHumiditiesWithLimit",
		strings.ToUpper("Get"),
		"/v1/humidities/{limit}",
		api.GetLastHumiditiesWithLimit,
	},

	Route{
		"GetCurrentPressure",
		strings.ToUpper("Get"),
		"/v1/pressure",
		api.GetCurrentPressure,
	},

	Route{
		"GetLastPressuresWithLimit",
		strings.ToUpper("Get"),
		"/v1/pressures/{limit}",
		api.GetLastPressuresWithLimit,
	},

	Route{
		"GetCurrentTemperature",
		strings.ToUpper("Get"),
		"/v1/temperature",
		api.GetCurrentTemperature,
	},

	Route{
		"GetLastTemperaturesWithLimit",
		strings.ToUpper("Get"),
		"/v1/temperatures/{limit}",
		api.GetLastTemperaturesWithLimit,
	},
}
