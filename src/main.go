package main

import (
	DB "github.com/dev-jpnobrega/api-rest/src/db"
	http "github.com/dev-jpnobrega/api-rest/src/infrastructure/helper"
	routers "github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
)

func main() {
	server := http.CreateEchoServer()

	database := &DB.DB{}
	database.Connect("postgres", "postgres://postgres:postgres@localhost/workout?sslmode=disable")

	routers.BuildRouters(server)

	server.Logger.Fatal(server.Start(":8080"))
}
