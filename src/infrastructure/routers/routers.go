package routers

import (
	"net/http"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
)

func adapter(command interfaces.ICommand, h handler.IHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r, command)
	}
}

// BuildRouters create router APP
func BuildRouters(server http.ServeMux) http.ServeMux {
	// router := mux.NewRouter().StrictSlash(true)

	server.HandleFunc(
		"/v1/people",
		adapter(
			factory.GetPeopleFactory(),
			handler.NewHandler(),
		),
	)

	return server
}
