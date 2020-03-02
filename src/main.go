package main

import (
	http "github.com/dev-jpnobrega/api-rest/src/infrastructure/helper"
	routers "github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
)

func main() {
	server := http.CreateEchoServer()

	routers.BuildRouters(server)

	server.Logger.Fatal(server.Start(":8080"))
}
