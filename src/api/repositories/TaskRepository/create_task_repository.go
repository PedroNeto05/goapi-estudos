package taskrepository

import "goApi/db/models"

func CreateTaskRepository(t *models.Task) error {
	err := db.Create(t).Error

	if err != nil {
		return err
	}
	return nil
}
