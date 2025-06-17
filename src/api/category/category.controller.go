package category

import "github.com/gofiber/fiber/v2"

type CategoryController interface {
	Create(c *fiber.Ctx) error
	ReadAll(c *fiber.Ctx) error
	Read(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
