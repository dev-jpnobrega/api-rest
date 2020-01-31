package infrastructure

import (
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	echo "github.com/labstack/echo/v4"
)

// IHandler ...
type IHandler interface {
	Handle(context echo.Context, command interfaces.ICommand) error
}
