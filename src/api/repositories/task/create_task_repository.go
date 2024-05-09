package taskrepository

import "goApi/db/models"

func CreateTaskRepository(task *models.Task) error {
	err := db.Create(task).Error

	if err != nil {
		return err
	}
	return nil
}
