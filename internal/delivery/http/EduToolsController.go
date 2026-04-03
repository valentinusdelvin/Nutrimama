package http

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"nutrimama/internal/model"
	"nutrimama/internal/usecase"
	"nutrimama/internal/utils"
	"strconv"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type EduToolsController struct {
	EduToolsUseCase *usecase.EduToolsUseCase
}

func NewEduToolsController(eduToolsUseCase *usecase.EduToolsUseCase) *EduToolsController {
	return &EduToolsController{
		EduToolsUseCase: eduToolsUseCase,
	}
}

func (c *EduToolsController) Create(ctx *fiber.Ctx) error {
	req := new(model.CreateEduToolsRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Dynamic Inline Image Uploader natively attached to Create Request mapping logic cleanly
	file, err := ctx.FormFile("thumbnail")
	if err == nil {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
			newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
			uploadDir := "./public/uploads"
			_ = os.MkdirAll(uploadDir, os.ModePerm)
			if ctx.SaveFile(file, fmt.Sprintf("%s/%s", uploadDir, newFileName)) == nil {
				req.Thumbnail = fmt.Sprintf("/uploads/%s", newFileName)
			}
		}
	}

	res, err := c.EduToolsUseCase.Create(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusCreated, "Request Successfull", res)
}

func (c *EduToolsController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}

	req := new(model.UpdateEduToolsRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	// Dynamic Inline Image Uploader natively attached mapping logic cleanly
	file, err := ctx.FormFile("thumbnail")
	if err == nil {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
			newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)
			uploadDir := "./public/uploads"
			_ = os.MkdirAll(uploadDir, os.ModePerm)
			if ctx.SaveFile(file, fmt.Sprintf("%s/%s", uploadDir, newFileName)) == nil {
				req.Thumbnail = fmt.Sprintf("/uploads/%s", newFileName)
			}
		}
	}
	req.ID = id

	res, err := c.EduToolsUseCase.Update(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) Get(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	if slug == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid dynamically mapped slug parameter"})
	}

	res, err := c.EduToolsUseCase.Get(slug)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) List(ctx *fiber.Ctx) error {
	res, err := c.EduToolsUseCase.List()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", res)
}

func (c *EduToolsController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id parameter"})
	}

	err = c.EduToolsUseCase.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return utils.SuccessResponse(ctx, fiber.StatusOK, "Request Successfull", nil)
}
