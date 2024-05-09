package taskusecase

import (
	taskrepository "goApi/api/repositories/TaskRepository"
	"goApi/db/models"
)

func GetTasksUseCase() []models.Task {
	tasks := taskrepository.GetTasksRepository()

	return tasks
}
