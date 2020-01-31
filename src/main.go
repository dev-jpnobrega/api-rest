package main

import (
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
	mid "github.com/labstack/echo/v4/middleware"
	echo "github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.Use(mid.Logger())
	server.Use(mid.Recover())

	routers.BuildRouters(server)

	server.Logger.Fatal(server.Start(":8080"))
}
