package users

import (
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	userController UserController
}

func NewUserRouter(userController UserController) *UserRouter {
	return &UserRouter{
		userController: userController,
	}
}

func (r *UserRouter) UserRoutes(router fiber.Router) {
	router.Get("/", r.userController.GetUserById)
	router.Put("/", r.userController.UpdateUser)
	router.Put("/password", r.userController.UpdateUserPassword)
}
