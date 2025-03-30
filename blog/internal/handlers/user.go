package handlers
import "github.com/gofiber/fiber/v2"
func GetAllUsers(c *fiber.Ctx) error   { return c.SendString("Get all") }
func GetUser(c *fiber.Ctx) error       { return c.SendString("Get one") }
func CreateUser(c *fiber.Ctx) error    { return c.SendString("Create") }
func UpdateUser(c *fiber.Ctx) error    { return c.SendString("Update") }
func DeleteUser(c *fiber.Ctx) error    { return c.SendString("Delete") }
