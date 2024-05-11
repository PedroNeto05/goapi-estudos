package taskrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func DeleteTaskRepository(taskID string, userID uuid.UUID) error {

	err := db.Where("id = ? AND  user_id = ?", taskID, userID).Delete(&models.Task{}).Error

	if err != nil {
		return err
	}

	return nil

}
