package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)
func main() {
    app := fiber.New()

    // Middleware
    app.Use(cors.New())
    app.Use(logger.New())

    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Start server
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}