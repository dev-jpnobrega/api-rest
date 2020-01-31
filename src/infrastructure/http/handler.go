package infrastructure

import (
	"fmt"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

// Handler ...
type Handler struct{}

// Handle ...
func (h *Handler) Handle(context echo.Context, c interfaces.ICommand) error {
	var model value.DataInput
	model.Args = context.QueryParams()
	// model.Authorization = r.Header.Get("Authorization")
	// model.XAppToken = r.Header.Get("x-app-token")

	err, rs := c.Execute(model)

	fmt.Println("Handler[result]", err, rs)

	return context.JSON(http.StatusOK, rs)
}

// NewHandler ...
func NewHandler() IHandler {
	return &Handler{}
}
