package http

import (
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type BookmarkController struct {
	BookmarkUseCase *usecase.BookmarkUseCase
}

func NewBookmarkController(uc *usecase.BookmarkUseCase) *BookmarkController {
	return &BookmarkController{BookmarkUseCase: uc}
}

func (c *BookmarkController) ToggleBookmark(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
	}

	// Read params reliably
	type ToggleRequest struct {
		EduToolsID int `json:"edu_tools_id"`
	}
	
	req := new(ToggleRequest)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	action, err := c.BookmarkUseCase.ToggleBookmark(userID, req.EduToolsID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	msg := "Successfully added bookmark"
	if action == "deleted" {
		msg = "Successfully removed bookmark"
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, msg, nil)
}

func (c *BookmarkController) GetBookmarks(ctx *fiber.Ctx) error {
	var userID int
	switch v := ctx.Locals("userID").(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid user id in token"})
	}

	bookmarks, err := c.BookmarkUseCase.GetUserBookmarks(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, "Successfully fetched user bookmarks", bookmarks)
}
