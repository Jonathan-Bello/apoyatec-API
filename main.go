package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// r := gin.Default()

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(200, "Hello, World!")
	// })

	// r.Run(":3000")

	port := os.Getenv("POST")

	if port == "" {
		// log.Fatal("$PORT must be set MENSO")
		port = "8080"
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"saludo": "holiwis",
		})
	})

	e.Start(":" + port)
}
