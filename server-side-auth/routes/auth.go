package routes

import (
	"fmt"
	"strings"

	"example.com/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewMiddleware() fiber.Handler {

	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()

	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}
	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}

type AuthUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	fmt.Println("Register")
	var data AuthUser

	bpErr := c.BodyParser(&data)
	if bpErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + bpErr.Error(),
		})
	}
	password, bcErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if bcErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + bcErr.Error(),
		})
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),
	}

	cuErr := models.CreateUser(&user)
	if cuErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + cuErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "registered",
	})
}

func Login(c *fiber.Ctx) error {
	var data AuthUser

	bpErr := c.BodyParser(&data)
	if bpErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + bpErr.Error(),
		})
	}

	var user models.User
	if !models.CheckEmail(data.Email, &user) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized - email not found",
		})
	}

	// password, bcErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	// if bcErr != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "something went wrong: " + bcErr.Error(),
	// 	})
	// }

	chapErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if chapErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized - password not found",
		})
	}

	sess, sgErr := store.Get(c)
	if sgErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + sgErr.Error(),
		})
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)

	ssErr := sess.Save()
	if ssErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + sgErr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged in",
	})
}

func Logout(c *fiber.Ctx) error {
	sess, sgErr := store.Get(c)
	if sgErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "no session: " + sgErr.Error(),
		})
	}
	dErr := sess.Destroy()
	if dErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + sgErr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}

func HealthCheck(c *fiber.Ctx) error {
	sess, sgErr := store.Get(c)
	if sgErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	auth := sess.Get(AUTH_KEY)

	if auth == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "authenticated",
	})
}
