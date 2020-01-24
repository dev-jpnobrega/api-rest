package routers

import (
	"net/http"
	"net/url"
	"reflect"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// GetPeopleParam live
type GetPeopleParam struct {
	Name string
	Age  int
}

func fillStruct(data url.Values, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)

		if !val.IsValid() {
			panic(val.Type())
		}

		val.Set(reflect.ValueOf(v[0]))
	}
}

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
