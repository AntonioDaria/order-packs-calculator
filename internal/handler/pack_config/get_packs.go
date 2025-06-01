package pack_config

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *PackConfigHandler) GetPacks(c *fiber.Ctx) error {
	packs, err := h.Service.GetPackSizes(context.Background())
	if err != nil {
		h.Logger.Error().Err(err).Msg("Failed to fetch pack sizes")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch pack sizes"})
	}

	return c.Status(http.StatusOK).JSON(packs)
}
