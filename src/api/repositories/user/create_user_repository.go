package userrepository

import "goApi/db/models"

func CreateUserRepository(user *models.User) error {
	err := db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}
