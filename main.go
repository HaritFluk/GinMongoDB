package main

import (
	"GinMongoDB/configs"
	"GinMongoDB/routes"


	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// run database
	configs.ConnectDB()

	// route
	routes.UserRoute(router) // add router

	// listen serve port 6000
	router.Run("localhost:6000")
}