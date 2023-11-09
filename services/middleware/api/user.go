package api

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRouter initializes the API routes
func (api *MiddlewareAPI) SetupRouter(router *fiber.App) {
	router.Get("/", api.GetIndexHandler)
}

// GetIndexHandler handles GET requests to the /:param endpoint
func (api *MiddlewareAPI) GetIndexHandler(c *fiber.Ctx) error {
	param := c.Params("param")
	indexData, err := api.app.GetIndex(param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": indexData})
}
