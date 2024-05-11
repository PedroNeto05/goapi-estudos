package helpers

import (
	"errors"
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"

	"gorm.io/gorm"
)

func UserExistByEmail(email string) (bool, error) {

	userExist, err := userrepository.GetUserByEmail(email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return !(userExist == (&models.User{})), nil
}
