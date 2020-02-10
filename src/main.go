package main

import (
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
	echo "github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return err
	}

	return
}

func main() {
	server := echo.New()
	server.Binder = &CustomBinder{}

	server.Use(mid.Logger())
	server.Use(mid.Recover())

	routers.BuildRouters(server)

	server.Logger.Fatal(server.Start(":8080"))
}
