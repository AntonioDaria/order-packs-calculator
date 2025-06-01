package pack_calculator

import (
	"sort"

	pack_calculator_service "github.com/AntonioDaria/order-packs-calculator/internal/service/pack_calculator"
	"github.com/gofiber/fiber/v2"
)

type CalculateRequest struct {
	Items int `json:"items"`
}

type CalculateResponse struct {
	Total int    `json:"total"`
	Packs []Pack `json:"packs"`
}

func (h *PackCalculatorHandler) Calculate(c *fiber.Ctx) error {
	var req CalculateRequest
	if err := c.BodyParser(&req); err != nil || req.Items <= 0 {
		h.logger.Warn().Err(err).Int("items", req.Items).Msg("Invalid request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: must include a positive integer 'items'",
		})
	}

	result, total, err := h.service.CalculatePacks(c.Context(), req.Items)
	if err != nil {
		switch err {
		case pack_calculator_service.ErrInvalidItemCount:
			h.logger.Warn().Err(err).Int("items", req.Items).Msg("Invalid item count from client")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		case pack_calculator_service.ErrNoValidPackCombination:
			h.logger.Info().Err(err).Int("items", req.Items).Msg("No valid pack combination found")
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		default:
			h.logger.Error().Err(err).Int("items", req.Items).Msg("Unexpected error calculating packs")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}
	}

	packs := make([]Pack, 0, len(result))
	for size, quantity := range result {
		packs = append(packs, Pack{
			Size:     size,
			Quantity: quantity,
		})
	}

	sort.Slice(packs, func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	})

	h.logger.Info().
		Int("items", req.Items).
		Int("total", total).
		Interface("packs", packs).
		Msg("Pack calculation successful")

	return c.Status(fiber.StatusOK).JSON(CalculateResponse{
		Total: total,
		Packs: packs,
	})
}

type Pack struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}
