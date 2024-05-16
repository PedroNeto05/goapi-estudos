package taskusecase

import (
	"errors"
	"goApi/api/helpers"
	taskrepository "goApi/api/repositories/task"

	"github.com/google/uuid"
)

func DeleteTaskUseCase(taskID string, userID uuid.UUID) error {

	userExist, err := helpers.UserExistByID(userID)

	if err != nil {
		return err
	}

	if !userExist {
		return errors.New("the user does not exist")
	}

	err = taskrepository.DeleteTaskRepository(taskID, userID)

	if err != nil {
		return err
	}

	return nil
}
