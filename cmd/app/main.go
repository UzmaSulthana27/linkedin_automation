package main

import (
	"log"

	"github.com/joho/godotenv"
	"linkedin-automation/internal/browser"
)

func main() {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	browser.StartBrowser()
}
