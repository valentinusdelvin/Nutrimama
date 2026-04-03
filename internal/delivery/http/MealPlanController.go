package http

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MealPlanController struct {
	MealPlanUseCase *usecase.MealPlanUseCase
}

func NewMealPlanController(uc *usecase.MealPlanUseCase) *MealPlanController {
	return &MealPlanController{MealPlanUseCase: uc}
}

func (c *MealPlanController) GenerateAIMealPlan(ctx *fiber.Ctx) error {
	role, ok := ctx.Locals("role").(string)
	if !ok || role != "mother" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "only mothers can trigger AI meal plan generation dynamically"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id mapping token securely"})
	}

	logs, err := c.MealPlanUseCase.GenerateAIMealPlan(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "SUCCESS! Gemini generated dynamic meal schedule mappings optimally securely!", logs)
}

func (c *MealPlanController) GetMealPlan(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id mapping token securely"})
	}

	logs, err := c.MealPlanUseCase.GetLatestMealPlan(userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", logs)
}
