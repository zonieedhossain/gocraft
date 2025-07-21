package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    app := echo.New()

    // Middleware
    app.Use(middleware.CORS())
    app.Use(middleware.Logger())

    // Routes
    app.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    // Start server
    if err := app.Start(":3000"); err != nil {
        panic(err)
    }
}