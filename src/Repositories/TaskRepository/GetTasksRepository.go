package taskrepository

import (
	models "goApi/src/Models"
)

func GetTasksRepository() []models.Task {
	tasks := []models.Task{}
	db.Find(&tasks)

	return tasks
}
