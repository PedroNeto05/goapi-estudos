package userusecase

import (
	"errors"
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"
)

func CreateUserUseCase(user *models.User) error {

	err := isAValidUser(user)

	if err != nil {
		return errors.New(err.Error())
	}

	userExist, err := userrepository.GetUserByEmail(user.Email)

	if err != nil {
		return errors.New(err.Error())
	}

	if userExist != (&models.User{}) {
		return errors.New("A tarefa ja existe")
	}
	return nil
}

func isAValidUser(user *models.User) error {
	if user.Email == "" && user.Name == "" && user.LastName == "" {
		return errors.New("O body esta vazio ou mal formatado")
	}

	if user.Email == "" {
		return errors.New("O email é requerido")
	}

	if user.Name == "" {
		return errors.New("O nome é requerida")
	}

	return nil
}
