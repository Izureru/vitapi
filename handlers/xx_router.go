package handlers

import (
	"net/http"

	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	r := pat.New()

	r.Add("OPTIONS", "/v1/users", http.HandlerFunc(HomeHandler))
	r.Add("OPTIONS", "/v1/users/{id}", http.HandlerFunc(HomeHandler))

	r.Get("/v1/users", GetUserHandler)

	r.Post("/v1/users", PostUserHandler)

	r.Delete("/v1/users/{id}", DeleteUserHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	r.Get("/", HomeHandler)
	return r
}
