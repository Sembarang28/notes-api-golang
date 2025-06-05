package middleware

import (
	"notes-management-api/src/helpers"
	"notes-management-api/src/shared/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UserSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Authorization header is missing",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid authorization format",
			})
		}

		accessToken := parts[1]
		if accessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Access token is missing",
			})
		}

		claims, err := helpers.VerifyAccessToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse{
				Code:    fiber.StatusUnauthorized,
				Status:  false,
				Message: "Invalid access token",
			})
		}

		c.Locals("userId", claims.UserID)

		// If we reach here, the authorization header is valid
		return c.Next()
	}
}
