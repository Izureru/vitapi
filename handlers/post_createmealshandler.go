package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/schnapi/entities"
	"github.com/DigitalInnovation/schnapi/global"
)

func PostUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("PostCollectMessagesHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()

	var userData entities.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	if userData.Name == "" {
		log.Println("Bad Request: User name property is empty")
		http.Error(rw, "Error User name property must not be empty", http.StatusBadRequest)
		return
	}

	err = global.Dal.CreateUser(userData)
	if err != nil {
		log.Printf("Error, failed to create message %v\n", err)
		http.Error(rw, "Error creating message", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
}
