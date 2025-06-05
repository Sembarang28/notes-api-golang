package users

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GetUserById(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UpdateUserPassword(c *fiber.Ctx) error
}
