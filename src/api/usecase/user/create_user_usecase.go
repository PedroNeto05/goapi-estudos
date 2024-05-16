package userusecase

import (
	"errors"
	"goApi/api/helpers"
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"
)

func CreateUserUseCase(user *models.User) error {

	err := isAValidUser(user)

	if err != nil {
		return err
	}
	userAlreadyExist, err := helpers.UserExistByEmail(user.Email)

	if err != nil {
		return err
	}

	if userAlreadyExist {
		return errors.New("email already registered")
	}

	userrepository.CreateUserRepository(user)
	return nil
}

func isAValidUser(user *models.User) error {
	if user.Email == "" && user.FirstName == "" && user.LastName == "" {
		return errors.New("the body is empty or poorly formatted")
	}

	if user.Email == "" {
		return errors.New("the email is required")
	}

	if user.FirstName == "" {
		return errors.New("the name is required")
	}

	return nil
}
