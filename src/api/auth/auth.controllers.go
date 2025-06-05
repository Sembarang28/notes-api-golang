package auth

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	RefreshTokenWeb(c *fiber.Ctx) error
	RefreshTokenMobile(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}
