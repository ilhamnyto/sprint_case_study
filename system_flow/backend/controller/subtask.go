package controller

import (
	"net/http"
	"strconv"

	"github.com/ilhamnyto/sprint_case_study/entity"
	"github.com/ilhamnyto/sprint_case_study/repositories"
	"github.com/labstack/echo/v4"
)

type SubTaskController struct {
	repo repositories.InterfaceSubTaskRepository
}

func NewSubTaskController(repo repositories.InterfaceSubTaskRepository) *SubTaskController {
	return &SubTaskController{repo: repo}
}

func (p *SubTaskController) CreateSubTask(c echo.Context) error {
	req := entity.SubTask{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := p.repo.Create(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask created successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *SubTaskController) GetOngoingSubTask(c echo.Context) error {

	tasks, err := p.repo.GetOngoingTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, tasks)
}

func (p *SubTaskController) GetCompletedSubTask(c echo.Context) error {

	tasks, err := p.repo.GetCompletedTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, tasks)
}

func (p *SubTaskController) DeleteSubTask(c echo.Context) error {
	taskId, err := strconv.Atoi(c.Param("taskId"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := p.repo.DeleteTask(taskId); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask deleted successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *SubTaskController) UpdateSubTask(c echo.Context) error {
	req := entity.SubTask{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := p.repo.UpdateTask(req.ID, req.Title); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask updated successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *SubTaskController) CompleteSubTask(c echo.Context) error {
	req := entity.SubTask{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := p.repo.CompleteTask(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask updated successfully"}
	return c.JSON(http.StatusOK, resp)
}