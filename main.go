package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func getStudent(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}
