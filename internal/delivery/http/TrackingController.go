package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TrackingController struct {
	TrackingUseCase *usecase.TrackingUseCase
}

func NewTrackingController(tUseCase *usecase.TrackingUseCase) *TrackingController {
	return &TrackingController{
		TrackingUseCase: tUseCase,
	}
}

func (c *TrackingController) SubmitTracking(ctx *fiber.Ctx) error {
	// Look up the role and user ID like we did in EditProfile
	role, ok := ctx.Locals("role").(string)
	if !ok || role != "mother" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "only mothers can submit nutrition tracking"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
	}

	req := new(model.SubmitTrackingRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.TrackingUseCase.SubmitTracking(userID, *req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Request Successfull", nil)
}

func (c *TrackingController) GetScores(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
	}

	date := ctx.Query("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	scores, err := c.TrackingUseCase.GetScores(userID, date)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully fetched live tracking scores", scores)
}
