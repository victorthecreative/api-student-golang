package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/victorthecreative/api-student-golang/db"
	_ "github.com/victorthecreative/students-system/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

// @title Student API
// @version 1.0
// @description This is a sample server Student API
// @host localhost:8080
// @BasePath /
// @schemes http
func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()

	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/student/:id", api.getStudent)
	api.Echo.PUT("/student/:id", api.updateStudent)
	api.Echo.DELETE("/student/:id", api.deleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (api *API) Start() error {

	return api.Echo.Start(":8080")
}
