package controllers

import (
	"context"
	"fmt"
	"hotelbooking/configs"
	"hotelbooking/models"
	"hotelbooking/response"
	"hotelbooking/util"
	"net/http"
	"time"

	//jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var UserDetails models.UserData
var UserDataCollection = configs.GetCollection(configs.DB, "RegisteredUser")

func RegisterUser(c *fiber.Ctx) error {
	var requestBodyData map[string]string

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	if err := c.BodyParser(&requestBodyData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HB_Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	if requestBodyData["password"] != requestBodyData["passwordmatch"] {
		return c.Status(http.StatusUnauthorized).JSON(response.HB_Response{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"data": "password does not match"}})

	}
	userid_in_string := UserDetails.Id.Hex()
	fmt.Println(userid_in_string)
	PasswordHashed, _ := bcrypt.GenerateFromPassword([]byte(requestBodyData["password"]), 14)
	newUserDetails := models.UserData{
		Id:        primitive.NewObjectID(),
		UserId:    userid_in_string,
		Firstname: requestBodyData["firstname"],
		Lastname:  requestBodyData["lastname"],
		Email:     requestBodyData["email"],
		Password:  PasswordHashed,
	}
	result, errInsertNewUser := UserDataCollection.InsertOne(ctx, newUserDetails)

	if errInsertNewUser != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.HB_Response{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": errInsertNewUser.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(response.HB_Response{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result.InsertedID}})

}

func Login(c *fiber.Ctx) error {
	var Login_Data map[string]string
	var Logged_in_User models.UserData

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	if errparse_login_data := c.BodyParser(&Login_Data); errparse_login_data != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HB_Response{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": errparse_login_data.Error()}})

	}

	err_check_user := UserDataCollection.FindOne(ctx, bson.M{"email": Login_Data["email"]}).Decode(&Logged_in_User)
	if err_check_user != nil {
		return c.Status(http.StatusNotFound).JSON(response.HB_Response{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User is not registered"}})

	}

	if err_check_password := bcrypt.CompareHashAndPassword(UserDetails.Password, []byte(Login_Data["password"])); err_check_password != nil {
		return c.Status(http.StatusUnauthorized).JSON(response.HB_Response{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"data": "Password is not correct"}})
	}

	newToken := util.GenerateJwt(UserDetails.UserId)

	new_cookie := fiber.Cookie{
		Name:     "hb_login_jwt",
		Value:    newToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&new_cookie)

	return c.Status(http.StatusOK).JSON(response.HB_Response{Status: http.StatusOK, Message: "user logged in", Data: &fiber.Map{"data": Logged_in_User.Email}})

}
