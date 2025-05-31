package controllers

import (
	"errors"
	"notes-management-api/src/helpers"
	"notes-management-api/src/internal/auth/dto"
	"notes-management-api/src/internal/auth/services"
	"notes-management-api/src/shared/response"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (h *AuthControllerImpl) Register(c *fiber.Ctx) error {
	var reqData dto.UserRegistrationRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	err := h.AuthService.Register(&reqData)
	if err != nil {
		switch {
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

	return c.Status(fiber.StatusCreated).JSON(response.APIResponse{
		Code:    fiber.StatusCreated,
		Status:  true,
		Message: "User registered successfully",
	})
}

func (h *AuthControllerImpl) Login(c *fiber.Ctx) error {
	var reqData dto.UserLoginRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	resData, err := h.AuthService.Login(&reqData)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrValidation):
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
				Code:    fiber.StatusBadRequest,
				Status:  false,
				Message: "Validation error",
			})
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid email or password",
			})
		case errors.Is(err, helpers.ErrUnauthorized):
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid email or password",
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
		Message: "Login successful",
		Data:    resData,
	})
}
