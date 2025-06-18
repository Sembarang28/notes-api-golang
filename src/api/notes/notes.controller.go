package notes

import "github.com/gofiber/fiber/v2"

type NotesController interface {
	Create(c *fiber.Ctx) error
	ReadAll(c *fiber.Ctx) error
	ReadOne(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
