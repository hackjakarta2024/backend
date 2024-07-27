package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/service"
	"go.uber.org/zap"
)

type fypHandler struct {
	fypService service.FypService
	logger     *zap.Logger
}

type FypHandler interface {
	GetFyp(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
}

func NewFypHandler(fypService service.FypService, logger *zap.Logger) FypHandler {
	return &fypHandler{
		fypService: fypService,
		logger:     logger,
	}
}

func (h *fypHandler) GetFyp(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	fyp, err := h.fypService.GetFyp(uuid.MustParse(userID))
	if err != nil {
		h.logger.Error("Error getting FYP", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    fyp,
	})

}

func (h *fypHandler) Search(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	query := c.Query("query")

	searchResp, err := h.fypService.Search(userID, query)
	if err != nil {
		h.logger.Error("Error searching FYP", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    searchResp,
	})

}
