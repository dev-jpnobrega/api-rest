package routers

import (
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	echo "github.com/labstack/echo/v4"
)

func adapter(command interfaces.ICommand, h handler.IHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		return h.Handle(c, command)
	}
}

// BuildRouters create router APP
func BuildRouters(server *echo.Echo) {
	// routeGroup := server.Group("/v1")

	server.GET(
		"/v1/people",
		adapter(
			factory.GetPeopleFactory(),
			handler.NewHandler(),
		),
	)

	server.POST(
		"/v1/login",
		adapter(
			factory.UserLoginFactory(),
			handler.NewHandler(),
		),
	)
}
