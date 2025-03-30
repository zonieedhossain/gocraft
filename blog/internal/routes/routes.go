package routes

import (
	"github.com/zonieed/blog/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(app interface{}) {
	f := app.(*fiber.App)
	f.Get("/hello", handlers.HelloFiber)
}
