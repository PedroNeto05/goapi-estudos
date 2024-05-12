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

func isAValidTask(task *models.Task) error {
	if task.Title == "" && task.Description == "" {
		return errors.New("o body esta vazio ou mal formatado")
	}

	if task.Title == "" {
		return errors.New("o titlo é requerido")
	}

	if task.Description == "" {
		return errors.New("a descrição é requerida")
	}

	return nil
}
