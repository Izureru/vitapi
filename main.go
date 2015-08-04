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
	global.LoadConfig()

	log.Println("I am listening to port", os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

	http.ListenAndServe(fmt.Sprintf(":%v", global.Config.Port), nil)
}
