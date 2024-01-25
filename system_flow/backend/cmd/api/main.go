package main

import (
	"os"

	"github.com/ilhamnyto/sprint_case_study/config"
	"github.com/ilhamnyto/sprint_case_study/controller"
	"github.com/ilhamnyto/sprint_case_study/pkg/database"
	"github.com/ilhamnyto/sprint_case_study/repositories"
	"github.com/ilhamnyto/sprint_case_study/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig(".env")
	
	db := database.ConnectDB()

	e := echo.New()

	taskRepository := repositories.NewTaskRepository(db.DbSQL)
	taskController := controller.NewTaskController(taskRepository)
	routes.TaskRoutes(e, *taskController)

	subTaskRepository := repositories.NewSubTaskRepository(db.DbSQL)
	subTaskController := controller.NewSubTaskController(subTaskRepository)
	routes.SubTaskRoutes(e, *subTaskController)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Sprint Asia Test!")
	})

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}