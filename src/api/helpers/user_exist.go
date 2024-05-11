package helpers

import (
	userrepository "goApi/api/repositories/user"
	"goApi/db/models"
)

func UserExist(email string) (bool, error) {

	userExist, err := userrepository.GetUserByEmail(email)

	if err != nil {
		return false, err
	}

	return (userExist == (&models.User{})), nil
}
