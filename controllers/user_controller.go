package controllers

import (
	"context"
	"net/http"
	"time"

	"GinMongoDB/configs"
	"GinMongoDB/models"
	"GinMongoDB/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollections *mongo.Collection = configs.GetCollection(configs.DB, "user")
var validator = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()
	
		// validate the request body
		if err := c.BindJSONRequest(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse(Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}))
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			
		}
	}
}