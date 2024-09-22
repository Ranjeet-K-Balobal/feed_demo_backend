package routes

import (
    "github.com/gofiber/fiber/v2"
    "feed_demo_backend/handlers" 
)

// Setup function to initialize all routes
func Setup(app *fiber.App) {
    // Root route
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the Feed Demo API!")
    })

    // Group API routes
    api := app.Group("/api") // Group for all /api routes

    // Group for user interactions
    userGroup := api.Group("/users")
    userGroup.Get("/", handlers.GetUserInteractions) // Define this in handlers
    userGroup.Post("/", handlers.CreateUserInteraction)

    // Other API groups can be created here
    // exampleGroup := api.Group("/example")
    // exampleGroup.Get("/", handlers.ExampleHandler)
}
