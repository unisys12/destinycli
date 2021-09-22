package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Eye's up guardian!")

	// Environment Setup should probably be moved to an Init method
	// or project scaffold. Just thinking outload.
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BUNKEY := os.Getenv("X-API-KEY")

	message := fmt.Sprintf("Using API Key: %v", BUNKEY)

	fmt.Println(message)
}
