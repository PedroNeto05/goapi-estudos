package helpers

import (
	"errors"
	"goApi/db/models"
)

func IsAValidTask(task *models.Task) error {
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
