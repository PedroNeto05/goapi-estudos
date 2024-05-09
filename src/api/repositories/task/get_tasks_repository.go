package taskrepository

import "goApi/db/models"

func GetTasksRepository() []models.Task {
	tasks := []models.Task{}
	db.Find(&tasks)

	return tasks
}
