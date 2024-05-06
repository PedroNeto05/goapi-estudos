package Router

import (
	taskcontrollers "goApi/src/Controllers/TaskControllers"

	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {

	v1 := router.Group("/api/v1")
	tasks := v1.Group("/tasks")
	taskcontrollers.InitializeTaskController()
	{
		// CREATE TASKS
		tasks.Post("/", taskcontrollers.CreateTaskController)

		// SHOW TASKS
		tasks.Get("/:taskId", taskcontrollers.GetTaskController)

		// INDEX TASKS
		tasks.Get("/", taskcontrollers.GetTasksController)

		// UPDATE TASKS
		tasks.Put("/", taskcontrollers.UpdateTaskController)

		// DELETE TASKS
		tasks.Delete("/", taskcontrollers.DeleteTaskController)
	}

}
