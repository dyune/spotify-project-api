package main

import (
	"log"
	"spotify-api/internal/api"
	"spotify-api/internal/handler"
)

func main() {
	app := handler.StartApplication()
	api.TestRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
