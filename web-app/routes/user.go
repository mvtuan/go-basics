package routes

import (
	"basic/web-app/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetAUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.EditUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
