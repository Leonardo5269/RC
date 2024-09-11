package main

import (
	"log"
	"os"

	"github.com/Leonardo5269/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func startApp() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")
	app.Static("/images", "./public/images")
	routes.SetupWebRoutes(app)
	routes.SetupApiRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func main() {
	startApp()
}
