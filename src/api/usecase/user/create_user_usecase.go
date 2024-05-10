package userusecase

import (
	"errors"
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"

	"gorm.io/gorm"
)

func CreateUserUseCase(user *models.User) error {

	err := isAValidUser(user)

	if err != nil {
		return err
	}
	userAlreadyExist, err := userrepository.GetUserByEmail(user.Email)

	if !(errors.Is(err, gorm.ErrRecordNotFound)) && err != nil {
		return err
	}

	if userAlreadyExist != nil {
		return errors.New("o usuario já existe")
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
