package taskrepository

import (
	models "goApi/src/Models"
)

func CreateTaskRepository(t *models.Task) error {
	err := db.Create(t).Error

	if err != nil {
		return err
	}
	return nil
}
