package taskusecase

import (
	"fmt"
	models "goApi/src/Models"
	taskrepository "goApi/src/Repositories/TaskRepository"
)

func CreateTaskUseCase(task *models.Task) error {

	if task.Title == "" && task.Description == "" {
		return fmt.Errorf("o body esta vazio ou mal formatado")
	}

	if task.Title == "" {
		return fmt.Errorf("o titlo é requerido")
	}

	if task.Description == "" {
		return fmt.Errorf("a descrição é requerida")
	}

	taskrepository.CreateTaskRepository(task)

	return nil
}
