package router

import (
	taskhandlers "goApi/api/handlers/task"
	userhandler "goApi/api/handlers/user"
	"goApi/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func initializeRoutes(router *fiber.App) {

	v1 := router.Group("/api/v1")
	users := v1.Group("/user")
	userhandler.InitializeUserHandler()
	{
		// CREATE USER
		users.Post("/", userhandler.CreateUserHandler)

		// LOGIN USER
		users.Post("/login", userhandler.LoginUserHandler)

		// // SHOW USER
		// users.Get("/:taskId")

		// // INDEX USER
		// users.Get("/")

		// // UPDATE USER
		// users.Put("/")

		// // DELETE USER
		// users.Delete("/")
	}
	tasks := v1.Group("/tasks")
	taskhandlers.InitializeTaskHandler()
	{
		// CREATE TASKS
		tasks.Post("/", middlewares.UserAuthenticated, taskhandlers.CreateTaskHandler)

		// SHOW TASKS
		// tasks.Get("/:taskId", middlewares.UserAuthenticated,taskhandlers.GetTaskHandler)

		// INDEX TASKS
		tasks.Get("/", middlewares.UserAuthenticated, taskhandlers.GetTasksHandler)

		// UPDATE TASKS
		tasks.Put("/", middlewares.UserAuthenticated, taskhandlers.UpdateTaskHandler)

		// DELETE TASKS
		tasks.Delete("/", middlewares.UserAuthenticated, taskhandlers.DeleteTaskHandler)
	}

}
