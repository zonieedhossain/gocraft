package handlers
import "github.com/gofiber/fiber/v2"

func HelloFiber(c *fiber.Ctx) error {
	return c.SendString("👋 Hello from Fiber handler!")
}
