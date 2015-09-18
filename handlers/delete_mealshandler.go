package handlers

import (
	"log"
	"net/http"

	"github.com/DigitalInnovation/schnapi/global"
)

func DeleteUserHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("DeleteUsersHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get(":id")

	err, errCode := global.Dal.DeleteUser(id)
	if err != nil {
		log.Printf("Error, failed to get meal Data %v\n", err)
		http.Error(rw, err.Error(), errCode)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	rw.WriteHeader(http.StatusNoContent)
}
