package main

import (
    "github.com/gofiber/fiber/v2"
)

type ScanImageRequest struct {
    ImageDataUri string `json:"imageDataUri"`
}

type ScanImageResponse struct {
    ShadeNames []string `json:"shadeNames"`
}

func main() {
    app := fiber.New()

    app.Post("/scan", func(c *fiber.Ctx) error {
        var req ScanImageRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        // Dummy logic: always return the same shades for demo
        resp := ScanImageResponse{
            ShadeNames: []string{"5.0 Light Brown", "4C Dark Chocolate"},
        }
        return c.JSON(resp)
    })

    app.Listen(":8080")
} 
