package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DigitalInnovation/simplyactiveapi/global"
)

func main() {

	router := NewRouter()

	log.Println("I am listening to port", os.Getenv("PORT"))

	// log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
	global.LoadConfig()
	log.Println(global.Config.Port)

	http.ListenAndServe(fmt.Sprintf(":%v", global.Config.Port), router)
}
