package main

import (
	"github.com/Harsh-apk/simpleBackendParamsTest/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:name", handlers.MainParamsHandler)
	app.Listen(":3000")

}
