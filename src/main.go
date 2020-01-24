package main

import (
	"context"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
	"log"
	"net/http"
)

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		ctx := context.WithValue(r.Context(), "Username", "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	server := http.NewServeMux()

	log.Fatal(http.ListenAndServe(":8080", routers.BuildRouters(server)))
}
