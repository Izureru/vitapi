package handlers

import (
	"log"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Home Handler Called")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
