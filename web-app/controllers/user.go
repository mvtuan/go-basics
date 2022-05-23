package controllers

import (
	"basic/web-app/common"
	"basic/web-app/configs"
	"basic/web-app/models"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func GetAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var user models.User
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	return c.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "success", Data: user})
}

func GetUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		}

		users = append(users, singleUser)
	}

	return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusOK, Message: "success", Data: users})
}

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var user models.User
	defer cancel()

	// validate request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, common.APIResponse{Status: http.StatusBadRequest, Message: "error"})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, common.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
	}

	var newUser = models.User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
	}

	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "Internal server error"})
	}
	return c.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "OK", Data: newUser})
}

func EditUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var user models.User

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, common.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, common.APIResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
	}

	update := bson.M{"first_name": user.FirstName, "last_name": user.LastName, "phone": user.Phone}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "success", Data: updatedUser})
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, common.APIResponse{Status: http.StatusNotFound, Message: "error", Data: "User with specified ID not found!"})
	}

	return c.JSON(http.StatusOK, common.APIResponse{Status: http.StatusOK, Message: "success", Data: "User successfully deleted!"})
}
