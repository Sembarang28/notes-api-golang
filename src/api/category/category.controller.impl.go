package category

import (
	"errors"
	"notes-management-api/src/api/category/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/shared/response"

	"github.com/gofiber/fiber/v2"
)

type CategoryControllerImpl struct {
	categoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) CategoryController {
	return &CategoryControllerImpl{
		categoryService: categoryService,
	}
}

func (h CategoryControllerImpl) Create(c *fiber.Ctx) error {
	var reqData dto.CategoryRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	userId := c.Locals("userId").(string)

	err := h.categoryService.Create(&reqData, userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrValidation):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Validation error",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(response.APIResponse{
		Code:    fiber.StatusCreated,
		Status:  true,
		Message: "Category created successfully",
	})
}

func (h CategoryControllerImpl) ReadAll(c *fiber.Ctx) error {
	search := c.Query("search")
	userId := c.Locals("userId").(string)

	categories, err := h.categoryService.ReadAll(search, userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Categories found",
		Data:    categories,
	})
}

func (h CategoryControllerImpl) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Locals("userId").(string)

	var reqData dto.CategoryRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	category, err := h.categoryService.Read(id, userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Category found",
		Data:    category,
	})
}

func (h CategoryControllerImpl) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Locals("userId").(string)

	var reqData dto.CategoryRequest
	err := h.categoryService.Update(&reqData, id, userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrValidation):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Validation error",
			})
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Category updated successfully",
	})
}

func (h CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Locals("userId").(string)

	err := h.categoryService.Delete(id, userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Category deleted successfully",
	})
}
