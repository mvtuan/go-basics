package main

import (
	"basic/web-app/configs"
	"basic/web-app/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &echo.Map{"data": "Hello from Echo & mongoDB"})
	})
	configs.ConnectDB()
	routes.UserRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
