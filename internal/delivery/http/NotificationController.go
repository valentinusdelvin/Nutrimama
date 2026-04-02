package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	NotificationUseCase *usecase.NotificationUseCase
}

func NewNotificationController(nUseCase *usecase.NotificationUseCase) *NotificationController {
	return &NotificationController{
		NotificationUseCase: nUseCase,
	}
}

func (c *NotificationController) Subscribe(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user token"})
	}

	req := new(model.SubscribePushRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.NotificationUseCase.Subscribe(userID, *req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully subscribed device to push notifications", nil)
}

func (c *NotificationController) GetList(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user token"})
	}

	notifications, err := c.NotificationUseCase.GetMyNotifications(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", notifications)
}
