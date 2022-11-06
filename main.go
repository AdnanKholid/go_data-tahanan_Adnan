package main

import (
	"mini_project/app/routes"
	"mini_project/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()

	server := echo.New()

	routes.RoutService(server)

	server.Logger.Fatal(server.Start(":6000"))
}
