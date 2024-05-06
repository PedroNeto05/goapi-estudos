package taskrepository

import (
	models "goApi/src/Models"
)

func CreateTaskRepository(t *models.Task) {
	db.Create(t)
}
