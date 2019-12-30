package main

import (
	"log"
	"net/http"

	"github.com/dev-jpnobrega/api-rest/src/infrastructure/router"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}

	http.Handle("/", s)

	log.Fatal(http.ListenAndServe(":8080", router.BuildRouters()))
}
