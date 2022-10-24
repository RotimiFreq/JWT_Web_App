package authmiddleware

import (
	response "hotelbooking/response"
	"hotelbooking/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Is_User_Authenticated(c *fiber.Ctx) error {
	current_cookie := c.Cookies("hb_login_jwt")

	if _, err_parse_jwt := util.ParseJWT(current_cookie); err_parse_jwt != nil {
		return c.Status(http.StatusUnauthorized).JSON(response.HB_Response{Status: http.StatusUnauthorized, Message: "error", Data: &fiber.Map{"data": "User is not Authenticated"}})
	}

	return c.Next()

}
