package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melhadou/hotel_booking/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"Users:": "Users"})
}

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Mohamed",
		LastName:  "Elhadouchi",
	}
	return c.JSON(u)
}
