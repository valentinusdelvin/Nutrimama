package http

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (c *UploadController) UploadImage(ctx *fiber.Ctx) error {
	// 1. Hook the raw Multipart Form File precisely!
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed grabbing binary image. Ensure your payload uses multipart/form-data with key 'image'",
		})
	}

	// 2. Validate Image Extension safely dodging malicious executable injection mathematically
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".webp" {
		return ctx.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Only .png, .jpg, .jpeg, and .webp images properly supported natively!",
		})
	}

	// 3. Generate a Crypto-Secure strict UUID masking the physical file solidly dodging naming crashes!
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	uploadDir := "./public/uploads"
	
	// Ensure pure directory initialization safely
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Crashed dynamically establishing physical storage backend cleanly.",
		})
	}

	savePath := fmt.Sprintf("%s/%s", uploadDir, newFileName)

	// 4. Save Binary logically via Native Fiber hook
	if err := ctx.SaveFile(file, savePath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Crashed dumping physical stream structurally.",
		})
	}

	// 5. Structure the dynamic Endpoint URL Mapping securely!
	fileUrl := fmt.Sprintf("/uploads/%s", newFileName)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Successfully processed upload securely mapping logic!",
		"thumbnail": fileUrl,
	})
}
