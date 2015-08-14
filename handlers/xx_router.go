package handlers

import (
	"net/http"

	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	r := pat.New()

	r.Add("OPTIONS", "/v1/meals", http.HandlerFunc(HomeHandler))
	r.Add("OPTIONS", "/v1/meals/{id}", http.HandlerFunc(HomeHandler))

	r.Get("/v1/meals", GetMealsHandler)

	r.Post("/v1/meals", PostMealsHandler)

	r.Delete("/v1/meals/{id}", DeleteMealHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	r.Get("/", HomeHandler)
	return r
}
