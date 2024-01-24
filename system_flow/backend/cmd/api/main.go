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

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}