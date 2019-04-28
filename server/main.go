package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/handle", handlers.Handler)
	http.HandleFunc("/registration", handlers.Registration)
	log.Print("Start server port 8080...")
	log.Print("Waiting clients...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
