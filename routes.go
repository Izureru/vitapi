package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MealIndex",
		"GET",
		"/meals",
		MealIndex,
	},
	Route{
		"ShowMeal",
		"GET",
		"/suggestions/{stepCount}",
		ShowMeal,
	},
	Route{
		"MealCreate",
		"POST",
		"/meals",
		MealCreate,
	},
}
