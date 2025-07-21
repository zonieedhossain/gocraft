package internal

import (
	"{{ .ModuleName }}/config"
	"fmt"
{{- if eq .Framework "fiber" }}
	"github.com/gofiber/fiber/v2"
	"{{ .ModuleName }}/internal/routes"
{{- else if eq .Framework "echo" }}
	"github.com/labstack/echo/v4"
	"{{ .ModuleName }}/internal/routes"
{{- else if eq .Framework "gin" }}
	"github.com/gin-gonic/gin"
	"{{ .ModuleName }}/internal/routes"
{{- end }}
)

func StartServer(cfg *config.Config) {
{{- if eq .Framework "fiber" }}
	app := fiber.New()
	routes.Register(app)
	app.Listen(fmt.Sprintf(":%s", cfg.Port))
{{- else if eq .Framework "echo" }}
	e := echo.New()
	routes.Register(e)
	e.Start(fmt.Sprintf(":%s", cfg.Port))
{{- else if eq .Framework "gin" }}
	r := gin.Default()
	routes.Register(r)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
{{- end }}
}
