package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Clifftech123/golang-postgres-api/initializers"
)

// global variable for the server
var (
	server *gin.Engine
)

//  init function to load the environment variables
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// connect to the database
	initializers.ConnectDB(&config)

	server = gin.Default()
}

// main function to run the server
func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// connect to the database
	router := server.Group("/api")
	router.GET("/first", func(ctx *gin.Context) {
		message := "Welcome to Golang api"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// connect to the database
	log.Fatal(server.Run(":" + config.ServerPort))
}

