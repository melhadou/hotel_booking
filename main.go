package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/melhadou/hotel_booking/api"
)

func main() {
	// reading port from command line
	var ListenPort = flag.String("port", "3000", "Listen Port")
	flag.Parse()

	app := fiber.New()
	appv1 := app.Group("/api/v1")

	appv1.Get("/user", api.HandleGetUsers)
	appv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*ListenPort)
}
