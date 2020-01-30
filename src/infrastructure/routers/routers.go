package routers

import (
	"context"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	"github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	"log"
	"net/http"
)

func adapter(command interfaces.ICommand, h handler.IHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r, command)
	}
}

func addContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)

		contextBound := context.Background()

		ctx := context.WithValue(contextBound, "Authentication", "TOKEN")

		// ctx = context.WithValue(ctx, "UserInfo", "{ 'personName': 'JP', 'personUid': 'UUID' }")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BuildRouters create router APP
func BuildRouters(server *http.ServeMux) http.Handler {
	server.HandleFunc(
		"/v1/people",
		adapter(
			factory.GetPeopleFactory(),
			handler.NewHandler(),
		),
	)

	context := addContext(server)

	return context
}
