package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main(){
    // Create a new Gin router
    r := gin.Default()

    // Middleware
    r.Use(cors.Default())
    r.Use(logger.SetLogger())

    // Routes
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, World!")
    })

    // Start server
    if err := r.Run(":3000"); err != nil {
        panic(err)
    }
}