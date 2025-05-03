package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-project/internal/db"
	"rest-project/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()

	log.Println("Starting the application...")
	routes.SetupRoutes(r)

	log.Println("Application started on port 8080")
	r.Run(":8080")
}
