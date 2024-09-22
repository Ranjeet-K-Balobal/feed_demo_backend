package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
	"feed_demo_backend/routes"
)

func main() {
    app := fiber.New()

    // Register routes from routes.go
    routes.Setup(app)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}
