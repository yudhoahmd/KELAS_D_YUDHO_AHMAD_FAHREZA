package httphandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	jwtware "github.com/gofiber/jwt/v2"
)

func TryHTTPHandler(app *fiber.App, secret string) {
	// Add logic in here to complete the tests
	basicauthHandler := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "rakamin",
		},
	})

	jwtauthHandler := jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			err = c.SendStatus(fiber.StatusUnauthorized)
			return err
		},
	})

	app.Use("/internal", basicauthHandler).Get("/internal", func(c *fiber.Ctx) (err error) {
		c.SendString("Welcome to the internal of Rakamin!")
		return
	})

	app.Use("/admin", jwtauthHandler).Get("/admin", func(c *fiber.Ctx) (err error) {
		c.SendString("Welcome back, admin!")
		return
	})
}
