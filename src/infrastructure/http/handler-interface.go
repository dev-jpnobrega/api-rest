package infrastructure

import (
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"net/http"
)

// IHandler ...
type IHandler interface {
	Handle(w http.ResponseWriter, r *http.Request, c interfaces.ICommand)
}
