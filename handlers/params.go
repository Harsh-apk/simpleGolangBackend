package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MainParamsHandler(c *fiber.Ctx) error {
	param := c.Params("name")
	page := fmt.Sprintf("<h1>Hello %s &#128075, how're you doing today?  </h1>", param)
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(page)
}
