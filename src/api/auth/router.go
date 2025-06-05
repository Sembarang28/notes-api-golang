package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	authController AuthController
}

func NewAuthRouter(authController AuthController) *AuthRouter {
	return &AuthRouter{
		authController: authController,
	}
}

func (r *AuthRouter) AuthRoutes(router fiber.Router) {
	router.Post("/register", r.authController.Register)
	router.Post("/login", r.authController.Login)
	router.Post("/refresh/web", r.authController.RefreshTokenWeb)
	router.Post("/refresh/mobile", r.authController.RefreshTokenMobile)
	router.Post("/logout", r.authController.Logout)
}
