package routes

import "github.com/gofiber/fiber/v2"

var APIs = map[string]func(*fiber.Ctx) error{
	"/getOffer": getOffer,
}

func SetupApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/:action?", func(c *fiber.Ctx) error {
		if action, exist := APIs[c.Params("action")]; exist {
			return action(c)
		}
		return c.Status(404).SendString("You requested an inexistent API")
	})
}

func getOffer(c *fiber.Ctx) error {
	return nil
}
