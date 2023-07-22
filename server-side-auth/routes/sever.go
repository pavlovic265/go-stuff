package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func RoutesSetup() {
	app := fiber.New()

	sStore := session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, // This is for https
		Expiration: time.Hour + 5,
	})
	store = sStore

	app.Use(logger.New())
	app.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	app.Post("/auth/register", Register)
	app.Post("/auth/login", Login)
	app.Post("/auth/logout", Logout)

	app.Get("/auth/healthcheck", HealthCheck)

	app.Get("/user", GetUser)

	log.Fatal(app.Listen(":8085"))
}
