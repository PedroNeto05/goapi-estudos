package taskusecase

import (
	taskrepository "goApi/api/repositories/task"
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetTasksUseCase(userID uuid.UUID) []models.Task {
	tasks := taskrepository.GetTasksRepository(userID)

	return tasks
}
