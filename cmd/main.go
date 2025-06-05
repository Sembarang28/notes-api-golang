package main

import (
	"notes-management-api/src/api"
	"notes-management-api/src/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	db := config.NewDatabaseConnection()
	validate := validator.New()

	app.Static("/public", "./public")

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowCredentials: true,
	// }))

	api.App(app, db, validate)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
