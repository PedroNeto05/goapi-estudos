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
	userAlreadyExist, err := helpers.UserExist(user.Email)

	if err != nil {
		return err
	}

	if userAlreadyExist {
		return errors.New("usuario já existe")
	}

	userrepository.CreateUserRepository(user)
	return nil
}

func isAValidUser(user *models.User) error {
	if user.Email == "" && user.Name == "" && user.LastName == "" {
		return errors.New("o body esta vazio ou mal formatado")
	}

	if user.Email == "" {
		return errors.New("o email é requerido")
	}

	if user.Name == "" {
		return errors.New("o nome é requerida")
	}

	return nil
}
