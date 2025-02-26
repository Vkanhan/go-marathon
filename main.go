package main

import (
	"log"

	"github.com/Vkanhan/go-marathon/config"
	"github.com/Vkanhan/go-marathon/server"
)

func main() {
	log.Println("Starting Runners App")
	log.Println("Initializing configuration")
	config := config.InitConfig("server")

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializing HTTP server")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}

// Controller layer - entry point for data - accept and handle http requests and routing and authorization
// Service layer - business logic - how data will be created and changed - validation
// Repository layer - prepare queries to be executed
// Databaes layer - consist of database tech
