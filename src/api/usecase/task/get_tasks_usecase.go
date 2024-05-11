package taskusecase

import (
	taskrepository "goApi/api/repositories/task"
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetTasksUseCase(userID uuid.UUID) ([]models.Task, error) {
	tasks, err := taskrepository.GetTasksRepository(userID)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
