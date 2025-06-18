package notes

import (
	"errors"
	"notes-management-api/src/api/notes/dto"
	"notes-management-api/src/helpers"
	"notes-management-api/src/shared/response"

	"github.com/gofiber/fiber/v2"
)

type NotesControllerImpl struct {
	noteService NoteService
}

func NewNotesController(noteService NoteService) NotesController {
	return &NotesControllerImpl{
		noteService: noteService,
	}
}

func (h NotesControllerImpl) Create(c *fiber.Ctx) error {
	var reqData dto.NotesRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	userId := c.Locals("userId").(string)

	err := h.noteService.Create(&reqData, userId)

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
		Message: "Notes created successfully",
	})
}
func (h NotesControllerImpl) ReadAll(c *fiber.Ctx) error
func (h NotesControllerImpl) ReadOne(c *fiber.Ctx) error
func (h NotesControllerImpl) Update(c *fiber.Ctx) error
func (h NotesControllerImpl) Delete(c *fiber.Ctx) error
