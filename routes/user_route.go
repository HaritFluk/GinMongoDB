package routes

import (
	"GinMongoDB/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// All routes relate to user here!!
	router.POST("/user", controllers.CreateUser()) // CreateUser route

	router.GET("/user/:userId", controllers.GetAUser()) //  Get a user route
	
	router.PUT("/user/:userId", controllers.EditAUser()) // Edit user route

	router.DELETE("/user/:userId", controllers.DeleteAUser()) // Delete user route

	router.GET("/users", controllers.GetAllUsers()) // Get all users route
}