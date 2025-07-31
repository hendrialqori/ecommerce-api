package handler

import (
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	CategoryUseCase usecase.CategoryUseCase
	logger          *logrus.Logger
}

func NewCategoryHandler(categoryUseCase usecase.CategoryUseCase, logger *logrus.Logger) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: categoryUseCase,
		logger:          logger,
	}
}

func (h *CategoryHandler) FindAll(c *fiber.Ctx) error {
	categories, err := h.CategoryUseCase.FindAll(c.Context())
	if err != nil {
		h.logger.WithError(err).Error("Failed to find all categories")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[[]*model.CategoryResponse]{
		Data:       categories,
		Message:    "Successfully retrieved all categories",
		StatusCode: fiber.StatusOK,
	})
}

func (h *CategoryHandler) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_category"))
	if err != nil || id <= 0 {
		h.logger.WithError(err).Error("Invalid category ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	category, err := h.CategoryUseCase.FindById(c.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to find category by ID")
		return err

	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.CategoryResponse]{
		Data:       category,
		Message:    "Category found successfully",
		StatusCode: fiber.StatusOK,
	})
}
func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	req := &model.CreateCategoryRequest{}
	if err := c.BodyParser(req); err != nil {
		h.logger.WithError(err).Error("Error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.CategoryUseCase.Create(c.Context(), req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create category")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&model.WebResponse[*model.CategoryResponse]{
		Data:       res,
		Message:    "Category created successfully",
		StatusCode: fiber.StatusCreated,
	})
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_category"))
	if err != nil || id <= 0 {
		h.logger.WithError(err).Error("Invalid category ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	req := &model.UpdateCategoryRequest{}
	req.ID = uint(id)
	if err := c.BodyParser(req); err != nil {
		h.logger.WithError(err).Error("Error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	res, err := h.CategoryUseCase.Update(c.Context(), req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update category")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.CategoryResponse]{
		Data:       res,
		Message:    "Category updated successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_category"))
	if err != nil || id <= 0 {
		h.logger.WithError(err).Error("Invalid category ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	err = h.CategoryUseCase.Delete(c.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to delete category")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[any]{
		StatusCode: fiber.StatusOK,
		Message:    "Category deleted successfully",
		Data:       nil,
	})
}
