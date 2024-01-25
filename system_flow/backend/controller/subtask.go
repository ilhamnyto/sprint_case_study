package controller

import (
	"net/http"
	"strconv"
	"time"

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

	req.CreatedAt = time.Now()

	if err := p.repo.CreateSubTask(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask created successfully"}
	return c.JSON(http.StatusOK, resp)
}

func (p *SubTaskController) GetOngoingSubTask(c echo.Context) error {

	subtasks, err := p.repo.GetOngoingSubTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(subtasks) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(200, subtasks)
}

func (p *SubTaskController) GetCompletedSubTask(c echo.Context) error {

	subtasks, err := p.repo.GetCompletedSubTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(subtasks) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(200, subtasks)
}

func (p *SubTaskController) DeleteSubTask(c echo.Context) error {
	subTaskId, err := strconv.Atoi(c.Param("subTaskId"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := p.repo.DeleteSubTask(subTaskId); err != nil {
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

	if err := p.repo.UpdateSubTask(req.ID, req.Title, req.Deadline); err != nil {
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

	if err := p.repo.CompleteSubTask(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := map[string]string{"message": "subtask updated successfully"}
	return c.JSON(http.StatusOK, resp)
}