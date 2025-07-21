package routes

{{- if eq .Framework "fiber" }}
import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
{{- else if eq .Framework "echo" }}
import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
{{- else if eq .Framework "gin" }}
import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})
}
{{- end }}
