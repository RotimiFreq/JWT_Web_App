package controllers

import (
	"context"
	"hotelbooking/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

)


func getAllUser(c *fiber.Ctx) error {
	var allUser []models.UserData

	ctx , cancel := context.WithTimeout(context.Background() , 20*time.Second)

	defer cancel()
	findOptions := options.Find()
	cursor, err_get_all_user := UserDataCollection.Find(ctx, bson.D{{}} ,  findOptions)
	if err_get_all_user != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.HB_Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err_get_all_user.Error()}})
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var get_all_user []models.UserData
		if errDecode := cursor.Decode(&get_all_user); errDecode != nil {
			return c.Status(http.StatusInternalServerError).JSON(response.HB_Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errDecode.Error()}})
		}

		allUser = append(allUser ,get_all_user )
	}

}

func Edit_User_Details(c *fiber.Ctx) error {
	var new_user_Details map[string]string

	var updated_user_details models.UserData

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	updated_user, err_updated_user := User
}