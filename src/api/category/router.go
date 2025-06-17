package category

import "github.com/gofiber/fiber/v2"

type CategoryRouter struct {
	categoryController CategoryController
}

func NewCategoryRouter(categoryController CategoryController) *CategoryRouter {
	return &CategoryRouter{
		categoryController: categoryController,
	}
}

func (r CategoryRouter) CategoryRouter(router fiber.Router) {
	router.Post("/", r.categoryController.Create)
	router.Get("/", r.categoryController.ReadAll)
	router.Get("/:id", r.categoryController.Read)
	router.Put("/:id", r.categoryController.Update)
	router.Delete("/:id", r.categoryController.Delete)
}
