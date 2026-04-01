package http

import (
	"strconv"
	
	"github.com/testuser/myapp-test/internal/domain"
	"github.com/testuser/myapp-test/internal/usecase"
	
	"github.com/gofiber/fiber/v2"
	
)

type Handler struct {
	logic *usecase.UserUsecase
}

func StartServer(logic *usecase.UserUsecase, port string) error {
	h := &Handler{logic: logic}

	
	app := fiber.New()
	api := app.Group("/api")
	api.Get("/users", h.GetAllUsers)
	api.Get("/users/:id", h.GetUserByID)
	api.Post("/users", h.CreateUser)
	return app.Listen(":" + port)

	
}


func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.logic.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.logic.GetUserByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.logic.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(user)
}


