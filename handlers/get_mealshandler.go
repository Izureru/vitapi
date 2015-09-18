package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/schnapi/entities"
	"github.com/DigitalInnovation/schnapi/global"
)

func GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetCollectMealsHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	var usersRetrieved []entities.User
	usersRetrieved, err = global.Dal.GetAllUsers()
	log.Println("GetAllMessages")

	if err != nil {
		log.Printf("Error, failed to get message Data %v\n", err)
		http.Error(rw, "Error getting message data", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(rw).Encode(usersRetrieved)
	if err != nil {
		log.Printf("Error, failed to encode JSON %v\n", err)
		http.Error(rw, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
