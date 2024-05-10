package userrepository

import (
	"goApi/db/models"
)

func GetUserByEmail(email string) (*models.User, error) {

	var user models.User
	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
