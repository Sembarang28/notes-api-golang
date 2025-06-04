package api

import (
	"notes-management-api/src/api/auth"
	"notes-management-api/src/api/auth/controllers"
	"notes-management-api/src/api/auth/repository"
	"notes-management-api/src/api/auth/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func App(app *fiber.App, db *gorm.DB, validate *validator.Validate) {
	api := app.Group("/api/v1")
	authRepositoryImpl := repository.NewAuthRepository(db)
	authServiceImpl := services.NewAuthService(authRepositoryImpl, validate)
	authControllerImpl := controllers.NewAuthController(authServiceImpl)
	authRouter := auth.NewAuthRouter(authControllerImpl)
	authRouter.AuthRoutes(api.Group("/auth"))

}
