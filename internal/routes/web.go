package routes

import "github.com/gofiber/fiber/v2"

var routes = map[string]string{
	"/":          "landing",
	"/contacts":  "contacts",
	"/solutions": "solutions",
}

func SetupWebRoutes(app *fiber.App) {
	app.Get("/:path?", func(c *fiber.Ctx) error {
		if tmpl, exist := routes[c.Params("path")]; exist {
			return c.Render(tmpl, nil)
		}
		return c.Status(404).Render("404", nil)
	})
}
