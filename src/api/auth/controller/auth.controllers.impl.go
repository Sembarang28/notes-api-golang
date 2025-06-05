package controller

import (
	"errors"
	"notes-management-api/src/api/auth"
	"notes-management-api/src/api/auth/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/shared/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) auth.AuthController {
	return &AuthControllerImpl{
		authService: authService,
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

	err := h.authService.Register(&reqData)
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

	resData, err := h.authService.Login(&reqData)
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

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    resData.RefreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Login successful",
		Data:    resData,
	})
}

func (h *AuthControllerImpl) RefreshTokenWeb(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  false,
			Message: "Refresh token is required",
		})
	}

	resData, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrUnauthorized):
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid refresh token",
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
		Message: "Token refreshed successfully",
		Data:    resData,
	})
}

func (h *AuthControllerImpl) RefreshTokenMobile(c *fiber.Ctx) error {
	var reqData dto.UserRefreshRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	resData, err := h.authService.RefreshToken(reqData.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrUnauthorized):
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid refresh token",
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
		Message: "Token refreshed successfully",
		Data:    resData,
	})
}

func (h *AuthControllerImpl) Logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
			Code:    fiber.StatusUnauthorized,
			Status:  false,
			Message: "Refresh token is required",
		})
	}

	err := h.authService.Logout(refreshToken)
	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrUnauthorized):
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid refresh token",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
				Code:    fiber.StatusInternalServerError,
				Status:  false,
				Message: "Internal server error",
			})
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(response.APIResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "Logout successful",
	})
}
