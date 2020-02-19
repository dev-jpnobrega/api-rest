package main

import (
	routers "github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
	"github.com/go-playground/validator"
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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	server := echo.New()
	server.Binder = &CustomBinder{}
	server.Validator = &CustomValidator{validator: validator.New()}

	// server.Use(mid.Logger())
	server.Use(mid.Recover())

	routers.BuildRouters(server)

	server.Logger.Fatal(server.Start(":8080"))
}
