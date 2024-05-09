package router

import (
	taskhandlers "goApi/api/handlers/task"

	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {

	v1 := router.Group("/api/v1")
	tasks := v1.Group("/tasks")
	taskhandlers.InitializeTaskHandler()
	{
		// CREATE TASKS
		tasks.Post("/", taskhandlers.CreateTaskHandler)

		// SHOW TASKS
		tasks.Get("/:taskId", taskhandlers.GetTaskHandler)

		// INDEX TASKS
		tasks.Get("/", taskhandlers.GetTasksHandler)

		// UPDATE TASKS
		tasks.Put("/", taskhandlers.UpdateTaskHandler)

		// DELETE TASKS
		tasks.Delete("/", taskhandlers.DeleteTaskHandler)
	}

}
