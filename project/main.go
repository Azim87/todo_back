package main

import (
	"todo/database"
	"todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	server := gin.Default()
	server.Use(gin.Recovery())
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
