package taskusecase

import (
	"errors"

	"goApi/api/helpers"
	taskrepository "goApi/api/repositories/task"
	"goApi/db/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTaskUseCase(task *models.Task, userID uuid.UUID) error {

	err := helpers.IsAValidTask(task)

	if err != nil {
		return err
	}

	taskAlreadyExist, err := taskrepository.GetTaskByTitleRepository(task.Title, userID)

	if !(errors.Is(err, gorm.ErrRecordNotFound)) && err != nil {
		return err
	}

	if taskAlreadyExist != nil {
		return errors.New("the task already exists")
	}

	err = taskrepository.CreateTaskRepository(task)

	if err != nil {
		return err
	}

	return nil
}
