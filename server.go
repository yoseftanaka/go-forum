package main

import (
	"forum/config"
	"forum/middlewares"
	"forum/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()
	config.InitRedis()
	e := echo.New()

	e.Use(middlewares.ErrorHandlerMiddleware)

	// Register all routes
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
