package users

import (
	"errors"
	"notes-management-api/src/api/users/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/shared/response"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	userService UserService
}

func NewUserController(userService UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (h *UserControllerImpl) GetUserById(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "User not found",
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
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func (h *UserControllerImpl) UpdateUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var userUpdateData dto.UserUpdateRequest
	if err := c.BodyParser(&userUpdateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	if err := h.userService.UpdateUser(userId, &userUpdateData); err != nil {
		switch {
		case errors.Is(err, helpers.ErrClientError):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Client error",
			})
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "User not found",
			})
		case errors.Is(err, helpers.ErrValidation):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Validation error",
			})
		case errors.Is(err, helpers.ErrEmailAlreadyExists):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(response.APIResponse{
				Code:    fiber.StatusUnprocessableEntity,
				Status:  false,
				Message: "Email already exists",
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
		Message: "User updated successfully",
	})
}

func (h *UserControllerImpl) UpdateUserPassword(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var userUpdatePasswordData dto.UserUpdatePasswordRequest
	if err := c.BodyParser(&userUpdatePasswordData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	if err := h.userService.UpdateUserPassword(userId, &userUpdatePasswordData); err != nil {
		switch {
		case errors.Is(err, helpers.ErrClientError):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Client error",
			})
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "User not found",
			})
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

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Password updated successfully",
	})
}
