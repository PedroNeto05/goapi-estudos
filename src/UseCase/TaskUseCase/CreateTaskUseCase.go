package taskusecase

import models "goApi/src/Models"

func CreateTaskUseCase(task models.Task) error {

	if task.Title == "" && task.Description == "" {

	}

	return nil
}
