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
func (h NotesControllerImpl) ReadAll(c *fiber.Ctx) error {
	var queries dto.NotesReadAllRequestQuery
	userId := c.Locals("userId").(string)
	if err := c.QueryParser(queries); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Status:  false,
			Message: "Internal server error",
		})
	}

	notes, err := h.noteService.ReadAll(queries.Search, userId, queries.Tags)

	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Notes data not found!",
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
		Message: "Notes data found!",
		Data:    notes,
	})
}

func (h NotesControllerImpl) ReadOne(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	id := c.Params("id")

	note, err := h.noteService.ReadOne(id, userId)

	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Note data not found!",
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
		Message: "Note data found!",
		Data:    note,
	})
}

func (h NotesControllerImpl) Update(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	id := c.Params("id")

	var reqData dto.NotesRequest
	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Invalid request body",
		})
	}

	err := h.noteService.Update(&reqData, id, userId)

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
				Message: "Note data not found!",
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
		Message: "Notes updated successfully",
	})
}

func (h NotesControllerImpl) Delete(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	id := c.Params("id")

	err := h.noteService.Delete(id, userId)

	if err != nil {
		switch {
		case errors.Is(err, helpers.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Code:    fiber.StatusNotFound,
				Status:  false,
				Message: "Note data not found!",
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
		Message: "Notes updated successfully",
	})
}
