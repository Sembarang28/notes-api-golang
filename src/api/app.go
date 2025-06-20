package api

import (
	"notes-management-api/src/api/auth"
	"notes-management-api/src/api/category"
	"notes-management-api/src/api/notes"
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

	categoryRepositoryImpl := category.NewCategoryRepository(db)
	categoryServiceImpl := category.NewCategoryService(categoryRepositoryImpl, validate)
	categoryControllerImpl := category.NewCategoryController(categoryServiceImpl)
	categoryRouter := category.NewCategoryRouter(categoryControllerImpl)
	categoryRouter.CategoryRouter(api.Group("/category", middleware.UserSession()))

	notesRepositoryImpl := notes.NewNotesRepository(db)
	notesServiceImpl := notes.NewNoteService(validate, notesRepositoryImpl)
	notesControllerImpl := notes.NewNotesController(notesServiceImpl)
	notesRouter := notes.NewNotesRouter(notesControllerImpl)
	notesRouter.NotesRouter(api.Group("/notes", middleware.UserSession()))
}
