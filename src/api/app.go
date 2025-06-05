package api

import (
	"notes-management-api/src/api/auth"
	"notes-management-api/src/api/users"
	"notes-management-api/src/shared/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func App(app *fiber.App, db *gorm.DB, validate *validator.Validate) {
	api := app.Group("/api/v1")
	authRepositoryImpl := auth.NewAuthRepository(db)
	authServiceImpl := auth.NewAuthService(authRepositoryImpl, validate)
	authControllerImpl := auth.NewAuthController(authServiceImpl)
	authRouter := auth.NewAuthRouter(authControllerImpl)
	authRouter.AuthRoutes(api.Group("/auth"))

	userRepositoryImpl := users.NewUserRepository(db)
	userServiceImpl := users.NewUserService(userRepositoryImpl, validate)
	userControllerImpl := users.NewUserController(userServiceImpl)
	userRouter := users.NewUserRouter(userControllerImpl)
	userRouter.UserRoutes(api.Group("/user", middleware.UserSession()))
}
