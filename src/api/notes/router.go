package notes

import "github.com/gofiber/fiber/v2"

type NotesRouter struct {
	notesController NotesController
}

func NewNotesRouter(notesController NotesController) *NotesRouter {
	return &NotesRouter{notesController: notesController}
}

func (r NotesRouter) NotesRouter(router fiber.Router) {
	router.Post("/", r.notesController.Create)
	router.Get("/", r.notesController.ReadAll)
	router.Get("/:id", r.notesController.ReadOne)
	router.Put("/:id", r.notesController.Update)
	router.Delete("/:id", r.notesController.Delete)
}
