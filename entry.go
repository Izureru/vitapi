package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DigitalInnovation/schnapi/global"
	"github.com/DigitalInnovation/schnapi/handlers"
)

func logEnvironmentVariables() {
	log.Printf("PORT: %v", os.Getenv("PORT"))
	log.Printf("MONGOURI: %v", os.Getenv("MONGOURI"))
	log.Printf("DBNAME: %v", os.Getenv("DBNAME"))
	log.Printf("CNAME: %v", os.Getenv("CNAME"))
}

func main() {

	logEnvironmentVariables()

	global.Setup()

	http.Handle("/", handlers.GetRouter())

	log.Println("Listening for connections on port ", global.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", global.Config.Port), nil)
}
