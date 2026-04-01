package http

import (
	"strconv"
	
	"net/http"
	
	"github.com/testuser/myapp-test/internal/domain"
	"github.com/testuser/myapp-test/internal/usecase"
	
	"github.com/gin-gonic/gin"
	
)

type Handler struct {
	logic *usecase.UserUsecase
}

func StartServer(logic *usecase.UserUsecase, port string) error {
	h := &Handler{logic: logic}

	
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/users", h.GetAllUsers)
	api.GET("/users/:id", h.GetUserByID)
	api.POST("/users", h.CreateUser)
	return r.Run(":" + port)
	
}


func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.logic.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.logic.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.logic.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, user)
}

