package infrastructure

import (
	"net/http"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	"github.com/go-playground/validator"
	echo "github.com/labstack/echo/v4"
)

type ErrorModel struct {
	Code    int
	Message string
}

type ResponseError struct {
	StatusCode int
	Errors     []ErrorModel
}

func onUnprocessableEntity(context echo.Context, err error) error {
	r := ResponseError{}

	r.StatusCode = http.StatusUnprocessableEntity

	for _, e := range err.(validator.ValidationErrors) {
		r.Errors = append(r.Errors, ErrorModel{
			Code:    2,
			Message: e.Field() + ".invalid",
		})
	}

	return context.JSONPretty(http.StatusUnprocessableEntity, r, " ")
}

func onSuccess(context echo.Context, r value.ResponseData) error {
	return context.JSONPretty(http.StatusOK, r, "  ")
}

// Handler ...
type Handler struct{}

// Handle ...
func (h *Handler) Handle(context echo.Context, c interfaces.ICommand) error {
	model := new(value.RequestData)
	model.Args = c.GetModelValidate()

	if err := context.Bind(&model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	if err := context.Validate(model.Args); err != nil {
		return onUnprocessableEntity(context, err)
	}

	_, rs := c.Execute(*model)

	return onSuccess(context, rs)
}

// NewHandler ...
func NewHandler() IHandler {
	return &Handler{}
}
