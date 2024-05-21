package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/Fodi999/text-style-service/color"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/styles/:classes", func(c *fiber.Ctx) error {
		classes := c.Params("classes")
		styles := color.ParseClasses(classes)
		var styleStrings []string
		for property, value := range styles {
			styleStrings = append(styleStrings, fmt.Sprintf("%s: %s;", property, value))
		}
		return c.SendString(strings.Join(styleStrings, " "))
	})

	log.Fatal(app.Listen(":3000"))
}
