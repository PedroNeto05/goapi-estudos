package taskusecase

import (
	models "goApi/src/Models"
	taskrepository "goApi/src/Repositories/TaskRepository"
)

func GetTasksUseCase() []models.Task {
	tasks := taskrepository.GetTasksRepository()

	return tasks
}
