package auth

import (
	"notes-management-api/src/api/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) *AuthRouter {
	return &AuthRouter{
		authController: authController,
	}
}

func (r *AuthRouter) AuthRoutes(router fiber.Router) {
	router.Post("/register", r.authController.Register)
	router.Post("/login", r.authController.Login)
}
