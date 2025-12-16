package main

import (
	"log"

	"linkedin-automation-poc/core"
	"linkedin-automation-poc/api"
	"linkedin-automation-poc/storage"
)

func main() {
	// Initialize SQLite DB
	db, err := storage.InitDB("linkedin.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Bot
	bot := core.NewBot(db)

	// Initialize API Server
	server := api.NewServer(bot, db)

	log.Println("Server running on http://localhost:8080")
	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
