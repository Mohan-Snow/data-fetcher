package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	// loads values from variables.env into the system
	if err := godotenv.Load("./configs/variables.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	githubUsername, exists := os.LookupEnv("BITGO_URL")
	if exists {
		fmt.Printf("Fetching data from %s", githubUsername)
	}
}
