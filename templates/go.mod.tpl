module {{.ModuleName}}

go {{.GoVersion}}

require (
	github.com/spf13/cobra v1.8.0
{{- if eq .Framework "fiber" }}
	github.com/gofiber/fiber/v2 v2.49.2
{{- else if eq .Framework "echo" }}
	github.com/labstack/echo/v4 v4.11.1
{{- else if eq .Framework "gin" }}
	github.com/gin-gonic/gin v1.9.1
{{- end }}
)
