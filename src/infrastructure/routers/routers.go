package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func handle(w http.ResponseWriter, r *http.Request, result value.DataResult) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.Write([]byte(result))

	json.NewEncoder(w).Encode(result)
}

var decoder = schema.NewDecoder()

var getPeopleParam struct {
	name string `schema:"name"`
	age  string `schema:"age"`
}

func adpater(command interfaces.ICommand) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var model value.DataInput

		model.Args = getPeopleParam

		keys := r.URL.Query()
		err2 := decoder.Decode(&model.Args, keys)

		fmt.Println("decode[GetPeopleParam][keys]", keys)
		fmt.Println("decode[GetPeopleParam][pars]", model.Args)
		fmt.Println("decode[err]", err2)

		// model.Args = keys
		err, p := command.Execute(model)

		fmt.Println("adpater[result]", err, p)

		handle(w, r, p)
	}
}

// BuildRouters create router APP
func BuildRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1/people", adpater(factory.GetPeopleFactory())).Methods("GET")

	return router
}
