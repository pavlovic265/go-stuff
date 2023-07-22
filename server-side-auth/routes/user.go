package routes

import (
	"fmt"

	"example.com/models"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	sess, sgErr := store.Get(c)
	if sgErr != nil {
		fmt.Println("sgErr :>> ", sgErr.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		fmt.Println("sess.Get(AUTH_KEY) == nil :>> ")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	userId := sess.Get(USER_ID)
	if userId == nil {
		fmt.Println("userId == nil :>> ")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// var user *models.User

	user, uErr := models.GetUser(fmt.Sprint(userId))
	if uErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + uErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
