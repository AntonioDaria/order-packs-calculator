package pack_config

import (
	"net/http"

	packconfig "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_config"
	"github.com/gofiber/fiber/v2"
)

type UpdatePacksResponse struct {
	Sizes []int `json:"sizes"`
}

func (h *PackConfigHandler) UpdatePacks(c *fiber.Ctx) error {
	var newPacks []int
	if err := c.BodyParser(&newPacks); err != nil {
		h.Logger.Warn().Err(err).Msg("Invalid pack sizes request body")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body — must be a JSON array of integers",
		})
	}

	err := h.Service.UpdatePackSizes(c.Context(), newPacks)
	if err != nil {
		switch err {
		case packconfig.ErrEmptyPackList, packconfig.ErrInvalidPackSize:
			h.Logger.Warn().Err(err).Msg("Validation error")
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		default:
			h.Logger.Error().Err(err).Msg("Repository failure during pack size update")
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}
	}

	return c.Status(http.StatusOK).JSON(
		UpdatePacksResponse{
			Sizes: newPacks,
		},
	)
}
