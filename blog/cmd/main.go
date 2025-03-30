package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zonieed/blog/internal/routes"
)

func main() {
	app := fiber.New()
	routes.Register(app)

	fmt.Println("Starting server on :8080")
	app.Listen(":8080")
}
