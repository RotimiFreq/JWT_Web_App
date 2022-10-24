package controllers

import (
	"context"

	"hotelbooking/models"
	"hotelbooking/response"
	"hotelbooking/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	bson "go.mongodb.org/mongo-driver/bson"
)

func Get_Current_User(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()
	var current_User models.UserData

	generated_JWT := c.Cookies("hb_login_jwt")

	current_user_token, _:= util.ParseJWT(generated_JWT)

	
	err_get_user := UserDataCollection.FindOne(ctx, bson.M{"user_id": current_user_token}).Decode(&current_User)
	if err_get_user != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.HB_Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": "can't find the user from token"}})

	}
	return c.Status(http.StatusOK).JSON(response.HB_Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": current_User}})

}

func Logout (c *fiber.Ctx) error {
	logout_cookie := fiber.Cookie{
		Name:     "hb_login_jwt",
		Value:    " ",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	
	c.Cookie(&logout_cookie)

	return c.Status(http.StatusOK).JSON(response.HB_Response{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User is logged out"}})
}


