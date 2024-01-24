package routes

import (
	"github.com/ilhamnyto/sprint_case_study/controller"
	"github.com/labstack/echo/v4"
)

func SubTaskRoutes(e *echo.Echo, c controller.SubTaskController) {
	subTaskGroup := e.Group("/api/v1/subtask")
	subTaskGroup.POST("/create", c.CreateSubTask)
	subTaskGroup.GET("/ongoing", c.GetOngoingSubTask)
	subTaskGroup.GET("/completed", c.GetCompletedSubTask)
	subTaskGroup.DELETE("/:taskId", c.DeleteSubTask)
	subTaskGroup.PUT("/update", c.UpdateSubTask)
	subTaskGroup.PUT("/complete", c.CompleteSubTask)
}