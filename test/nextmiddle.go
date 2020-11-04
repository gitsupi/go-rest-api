package main

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"radical.com/go-rest-api/db"
	"radical.com/go-rest-api/test/pkg"
)

var Collection = db.UserCollection

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, ss!")
	})
	pkg.NonAuthentically(app)
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Content-type", "application/json")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
		// Go to next middleware:
		return c.Next()
	})

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	pkg.Authentically(app, Collection)

	app.Listen(":9090")
}
