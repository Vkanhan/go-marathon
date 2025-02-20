package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}
	server := &http.Server{
		Addr:    ":" + portString,
		Handler: nil,
	}
	log.Printf("Listening to port: %v", portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
