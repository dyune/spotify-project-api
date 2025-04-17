package main

import (
	"log"
	"spotify-api/internal"
)

func main() {
	app := internal.StartApplication()
	internal.TestRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
