package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getPeople(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func BuildRouters() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/v1/people", getPeople).Methods("GET")

	return router
}
