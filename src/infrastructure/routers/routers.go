package infrastructure

import (
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	factory "github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	echo "github.com/labstack/echo/v4"
)

func Adapter(command interfaces.ICommand, h handler.IHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		return h.Handle(c, command)
	}
}

// BuildRouters create router APP
func BuildRouters(server *echo.Echo) {
	// routeGroup := server.Group("/v1")

	server.GET(
		"/v1/people",
		Adapter(
			factory.GetPeopleFactory(),
			handler.NewHandler(),
		),
	)

	server.POST(
		"/v1/login",
		Adapter(
			factory.UserLoginFactory(),
			handler.NewHandler(),
		),
	)
}
