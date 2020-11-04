
package main

import (
"github.com/gofiber/fiber"
"github.com/gofiber/template/html"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./test/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Listen(":3000")
}
