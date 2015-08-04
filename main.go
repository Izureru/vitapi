package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := GetPort()
	log.Println("[-] Listening on...", port)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "hello, world")
	})

	err := http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
	// router := NewRouter()

	// log.Println("I am listening to port", os.Getenv("PORT"))

	// log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

	// http.Handle("/", router)

	// http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return port
}
