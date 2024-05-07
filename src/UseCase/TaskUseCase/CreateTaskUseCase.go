package taskusecase

import (
	"fmt"
	models "goApi/src/Models"
	taskrepository "goApi/src/Repositories/TaskRepository"
)

func CreateTaskUseCase(task *models.Task) error {

	err := isAValidTask(task)

	if err != nil {
		return fmt.Errorf(" %v", err)
	}

	taskAlreadyExist, err := taskrepository.GetTaskByTitleRepository(task.Title)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	if taskAlreadyExist != (&models.Task{}) {
		return fmt.Errorf("a tarefa ja existe")
	}

	err = taskrepository.CreateTaskRepository(task)

	if err != nil {
		return fmt.Errorf("erro na criação da tarefa")
	}

	return nil
}

func isAValidTask(task *models.Task) error {
	if task.Title == "" && task.Description == "" {
		return fmt.Errorf("o body esta vazio ou mal formatado")
	}

	if task.Title == "" {
		return fmt.Errorf("o titlo é requerido")
	}

	if task.Description == "" {
		return fmt.Errorf("a descrição é requerida")
	}

	return nil
}
