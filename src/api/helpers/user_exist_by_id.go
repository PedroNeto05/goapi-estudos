package helpers

import (
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"

	"github.com/google/uuid"
)

func UserExistByID(id uuid.UUID) (bool, error) {

	userExist, err := userrepository.GetUserByID(id)

	if err != nil {
		return false, err
	}

	return !(userExist == (&models.User{})), nil
}
