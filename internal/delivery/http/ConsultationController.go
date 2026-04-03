package http

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ConsultationController struct {
	ConsultationUseCase *usecase.ConsultationUseCase
}

func NewConsultationController(uc *usecase.ConsultationUseCase) *ConsultationController {
	return &ConsultationController{ConsultationUseCase: uc}
}

func (c *ConsultationController) BookSession(ctx *fiber.Ctx) error {
	role, ok := ctx.Locals("role").(string)
	if !ok || role != "mother" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "only mothers can securely book consultations"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token mapping"})
	}

	req := new(usecase.BookSessionRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	session, err := c.ConsultationUseCase.BookSession(userID, *req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Successfully booked and purchased consultation schedule", session)
}

func (c *ConsultationController) EndSession(ctx *fiber.Ctx) error {
	role, ok := ctx.Locals("role").(string)
	if !ok {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
	}

	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token mapping"})
	}

	idStr := ctx.Params("id")
	consultationId, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid consultation id"})
	}

	session, err := c.ConsultationUseCase.EndSession(userID, role, consultationId)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully terminated consultation session securely", session)
}

func (c *ConsultationController) GetInbox(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token mapping"})
	}

	role, ok := ctx.Locals("role").(string)
	if !ok {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "role dynamically missing aggressively"})
	}

	inbox, err := c.ConsultationUseCase.GetInboxList(userID, role)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully retrieved native inbox structured mappings", inbox)
}

