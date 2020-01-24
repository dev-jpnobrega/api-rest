package infrastructure

import (
	"encoding/json"
	"fmt"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	"net/http"
)

// Handler ...
type Handler struct{}

// Handle ...
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request, c interfaces.ICommand) {
	var model value.DataInput
	model.Args = r.URL.Query()
	model.Authorization = r.Header.Get("Authorization")
	model.XAppToken = r.Header.Get("x-app-token")

	ctx := r.Context()

	err, rs := c.Execute(model)

	fmt.Println("Handler[result]", err, rs, ctx)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(rs)
}

// NewHandler ...
func NewHandler() IHandler {
	return &Handler{}
}
