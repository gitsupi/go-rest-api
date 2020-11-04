package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"time"
)

func main() {
	fmt.Printf("ddddddddddsss")

	app := fiber.New()

	// Login route
	app.Post("/login", login)
	app.Get("/login", accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)
	// Unauthenticated route
	app.Get("/", accessible)

	app.Listen(":3001")
}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	fmt.Printf("%v, %v", user, pass)

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessibled(c *fiber.Ctx) error {
	fmt.Printf("ddddddddddsss")

	return c.SendString("Accessible")
}

func accessible(c *fiber.Ctx) error {
	fmt.Printf("ddddddddddsss")

	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	fmt.Printf("in resiricted")

	user := c.Locals("user").(*jwt.Token)
	fmt.Printf("%v", user)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.SendString("Welcome " + name)
}
