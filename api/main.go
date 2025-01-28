package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"cryptolock/Connection"
	"cryptolock/Models"
	"cryptolock/Routes"
)

func main() {
	// Connect to the database
	Connection.Connect()

	// Migrate the User model
	err := Connection.DB.AutoMigrate(&Models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Set up the router
	router := gin.Default()
	Routes.SetupRoutes(router)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
