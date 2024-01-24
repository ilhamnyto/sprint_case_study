package routes

import (
	"github.com/ilhamnyto/sprint_case_study/controller"
	"github.com/labstack/echo/v4"
)

func TaskRoutes(e *echo.Echo, c controller.TaskController) {
	tasksGroup := e.Group("/api/v1/tasks")
	tasksGroup.POST("/create", c.CreateTask)
	tasksGroup.GET("/ongoing", c.GetOngoingTask)
	tasksGroup.GET("/completed", c.GetCompletedTask)
	tasksGroup.DELETE("/:taskId", c.DeleteTask)
	tasksGroup.PUT("/update", c.UpdateTask)
	tasksGroup.PUT("/complete", c.CompleteTask)
}