package taskusecase

import (
	taskrepository "goApi/api/repositories/task"
	"goApi/db/models"

	"github.com/google/uuid"
)

func UpdateTaskUseCase(task *models.Task, taskID string, userID uuid.UUID) error {

	err := taskrepository.UpdateTaskRepository(task, taskID, userID)

	if err != nil {
		return err
	}

	return nil
}
