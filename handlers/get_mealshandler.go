package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/simplyactiveapi/entities"
	"github.com/DigitalInnovation/simplyactiveapi/global"
)

func GetMealsHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetCollectMealsHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	unexpired := r.URL.Query().Get("unexpired")

	var mealsRetrieved []entities.Meal

	if unexpired != "" && unexpired == "true" {
		mealsRetrieved, err = global.Dal.GetAllMeals()
		log.Println("GetUnexpiredMessages")
	} else {
		mealsRetrieved, err = global.Dal.GetAllMeals()
		log.Println("GetAllMessages")
	}

	if err != nil {
		log.Printf("Error, failed to get message Data %v\n", err)
		http.Error(rw, "Error getting message data", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(rw).Encode(mealsRetrieved)
	if err != nil {
		log.Printf("Error, failed to encode JSON %v\n", err)
		http.Error(rw, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
