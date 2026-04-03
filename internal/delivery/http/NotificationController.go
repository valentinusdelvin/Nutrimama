package http

import (
	"nutrimama/internal/model"
	"nutrimama/internal/service"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	NotificationUseCase *usecase.NotificationUseCase
	PushService         *service.PushService
}

func NewNotificationController(nUseCase *usecase.NotificationUseCase, pushSvc *service.PushService) *NotificationController {
	return &NotificationController{
		NotificationUseCase: nUseCase,
		PushService:         pushSvc,
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

func (c *NotificationController) GetVapidPublicKey(ctx *fiber.Ctx) error {
	key := os.Getenv("VAPID_PUBLIC_KEY")
	if key == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "VAPID not configured"})
	}
	return ctx.JSON(fiber.Map{"public_key": key})
}

func (c *NotificationController) SendTest(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user token"})
	}

	req := new(model.SendTestPushRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.NotificationUseCase.SendTestNotification(userID, req.SubscriptionID, c.PushService); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully sent test notification", nil)
}
