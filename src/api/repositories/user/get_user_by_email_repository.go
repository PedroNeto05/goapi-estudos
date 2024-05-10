package userrepository

import (
	"fmt"
	"goApi/db/models"
)

func GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := db.Where("title = ?", email).First(&user).Error

	fmt.Println("PASSEI")

	if err != nil {
		return nil, err
	}

	return &user, nil
}
