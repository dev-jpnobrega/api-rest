package routers

import (
	"net/http"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	"github.com/gorilla/mux"
)

func adapter(command interfaces.ICommand, h handler.IHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r, command)
	}
}

// BuildRouters create router APP
func BuildRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/v1/people",
		adapter(
			factory.GetPeopleFactory(),
			handler.NewHandler(),
		),
	).Methods("GET")

	return router
}
