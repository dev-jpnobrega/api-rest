package infrastructure

import (
	"fmt"
	"net/http"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	echo "github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct{}

// Handle ...
func (h *Handler) Handle(context echo.Context, c interfaces.ICommand) error {
	model := new(value.DataInput)
	model.Args = c.GetModelValidate()

	if err := context.Bind(&model.Args); err != nil {
		fmt.Println("Handler[context.Bind]", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	err1, rs := c.Execute(*model)

	fmt.Println("Handler[result]", err1, rs)

	return context.JSONPretty(http.StatusOK, rs, "  ")
}

// NewHandler ...
func NewHandler() IHandler {
	return &Handler{}
}
