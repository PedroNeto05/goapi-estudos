package taskusecase

import (
	"errors"

	"goApi/api/helpers"
	taskrepository "goApi/api/repositories/task"
	"goApi/db/models"

	"gorm.io/gorm"
)

func CreateTaskUseCase(task *models.Task) error {

	err := helpers.IsAValidTask(task)

	if err != nil {
		return err
	}

	taskAlreadyExist, err := taskrepository.GetTaskByTitleRepository(task.Title)

	if !(errors.Is(err, gorm.ErrRecordNotFound)) && err != nil {
		return err
	}

	if taskAlreadyExist != nil {
		return errors.New("a tarefa já existe")
	}

	err = taskrepository.CreateTaskRepository(task)

	if err != nil {
		return err
	}

	return nil
}
