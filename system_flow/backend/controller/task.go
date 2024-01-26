package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ilhamnyto/sprint_case_study/entity"
	"github.com/ilhamnyto/sprint_case_study/repositories"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repo repositories.InterfaceTaskRepository
}

func NewTaskController(repo repositories.InterfaceTaskRepository) *TaskController {
	return &TaskController{repo: repo}
}

func (p *TaskController) CreateTask(c echo.Context) error {
	req := entity.Task{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.CreatedAt = time.Now()
	if req.Deadline != nil {
		parsedTime := time.Date(req.Deadline.Year(), req.Deadline.Month(), req.Deadline.Day(), 0, 0, 0, 0, time.UTC)
		req.Deadline = &parsedTime
	}

	task, err := p.repo.CreateTask(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]interface{}{"message": "task created successfully", "data": task}
	return c.JSON(http.StatusOK, resp)
}

func (p *TaskController) GetOngoingTask(c echo.Context) error {

	tasks, err := p.repo.GetOngoingTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(200, tasks)
}

func (p *TaskController) GetTaskAndSubTask(c echo.Context) error {

	tasks, err := p.repo.GetTaskAndSubTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(200, tasks)
}

func (p *TaskController) GetCompletedTask(c echo.Context) error {

	tasks, err := p.repo.GetCompletedTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(tasks) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(200, tasks)
}

func (p *TaskController) DeleteTask(c echo.Context) error {
	taskId, err := strconv.Atoi(c.Param("taskId"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := p.repo.DeleteTask(taskId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "task deleted successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *TaskController) UpdateTask(c echo.Context) error {
	req := entity.Task{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := p.repo.UpdateTask(req.ID, req.Title, req.Deadline); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "task updated successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *TaskController) CompleteTask(c echo.Context) error {
	req := entity.Task{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	complete := time.Now()
	req.CompletedAt = &complete

	if err := p.repo.CompleteTask(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "task updated successfully"}
	return c.JSON(http.StatusOK, resp)
}
