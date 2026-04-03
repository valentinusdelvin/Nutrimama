package http

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ConsultantController struct {
	ConsultantUseCase *usecase.ConsultantUseCase
}

func NewConsultantController(uc *usecase.ConsultantUseCase) *ConsultantController {
	return &ConsultantController{ConsultantUseCase: uc}
}

func (c *ConsultantController) List(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	search := ctx.Query("search", "")

	paginatedResult, err := c.ConsultantUseCase.ListConsultants(search, page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully fetched mapped consultants natively", paginatedResult)
}
