package userrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetUserByID(id uuid.UUID) (*models.User, error) {

	var user models.User
	err := db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
