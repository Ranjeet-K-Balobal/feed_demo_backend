package handlers

import "github.com/gofiber/fiber/v2"

// GetUserInteractions - example handler
func GetUserInteractions(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Get user interactions",
    })
}

// CreateUserInteraction - example handler
func CreateUserInteraction(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status":  "success",
        "message": "User interaction created",
    })
}
