package main

import (
	"fmt"
	"log"

	"github.com/Clifftech123/golang-postgres-api/initializers"
	"github.com/Clifftech123/golang-postgres-api/models"
)

// init function to load the environment variables
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

// main function to run the migrations
func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}

