package main

import (
	"log"
	"os"

	"github.com/Vkanhan/go-marathon/server"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetConfigFile("server.toml")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	dbHandler := server.InitDatabase(config)
	httpServer := server.InitHttpServer(config, dbHandler)

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	log.Printf("Listening to port: %v", portString)
	httpServer.Start()
}

// Controller layer - entry point for data - accept and handle http requests and routing and authorization 
// Service layer - business logic - how data will be created and changed - validation
// Repository layer - prepare queries to be executed
// Databaes layer - consist of database tech 