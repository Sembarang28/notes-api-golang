package controllers

import "notes-management-api/src/api/users/repository"

type UserControllerImpl struct {
	userController repository.UserRepository
}

func NewUserController(userController repository.UserRepository) UserController {
	return &UserControllerImpl{
		userController: userController,
	}
}
