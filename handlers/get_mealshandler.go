package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/vitamns/entities"
	"github.com/DigitalInnovation/vitamns/global"
)

func GetMealsHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetMealsHandler called")

	var mealsRetrieved []entities.Meal
	mealsRetrieved, err := global.Dal.GetAllMeals()
	log.Println("GetAllMessages")

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
