package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ilhamnyto/sprint_case_study/entity"
)

var (
	queryCreateTask = `
		INSERT INTO task (title, created_at, deadline) VALUES (?, ?, ?)
	`
	queryGetOngoingTask = `
		SELECT id, title, created_at, deadline, completed_at from task WHERE completed_at is null
	`
	queryDeleteTask = `
		DELETE FROM task WHERE id = ?
	`
	queryGetCompletedTask = `
		SELECT id, title, created_at, deadline, completed_at from task WHERE completed_at is not null
	`
	queryGetTaskAndSubTask = `
		SELECT t.*, ifnull(st.id, 0) as subtask_id, ifnull(st.task_id, 0) as task_id, ifnull(st.title, "") as subtask_title, st.deadline as subtask_deadline, 
		ifnull(t.created_at, DATE("2017-06-15")) as subtask_created_at, st.completed_at as subtask_completed_at from task as t left join subtask as st on t.id = st.task_id
		order by t.id DESC
	`
	queryUpdateTask = `
		UPDATE task SET title = ?, deadline = ? WHERE id = ?
	`
	queryCompleteTask = `
		UPDATE task SET completed_at = ? WHERE id = ?
	`
)

type InterfaceTaskRepository interface {
	CreateTask(task *entity.Task) error
	GetOngoingTask() ([]*entity.Task, error)
	GetTaskAndSubTask() ([]*entity.Task, error)
	GetCompletedTask() ([]*entity.Task, error)
	DeleteTask(id int) error
	UpdateTask(id int, title string, deadline *time.Time) error
	CompleteTask(id int) error
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) InterfaceTaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *entity.Task) error {
	stmt, err := r.db.Prepare(queryCreateTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(task.Title, task.CreatedAt, task.Deadline); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetOngoingTask() ([]*entity.Task, error) {
	stmt, err := r.db.Prepare(queryGetOngoingTask)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*entity.Task

	for rows.Next() {
		tempTask := new(entity.Task)
		if err := rows.Scan(&tempTask.ID, &tempTask.Title, &tempTask.CreatedAt, &tempTask.Deadline, &tempTask.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, tempTask)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskAndSubTask() ([]*entity.Task, error) {
	stmt, err := r.db.Prepare(queryGetTaskAndSubTask)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resTasks []*entity.Task

	for rows.Next() {
		tempTask := new(entity.Task)
		tempSubTask := new(entity.SubTask)
		if err := rows.Scan(&tempTask.ID, &tempTask.Title, &tempTask.CreatedAt, &tempTask.Deadline, &tempTask.CompletedAt, &tempSubTask.ID, &tempSubTask.TaskID, &tempSubTask.Title, &tempSubTask.Deadline, &tempSubTask.CreatedAt, &tempSubTask.CompletedAt); err != nil {
			fmt.Println(err)
			return nil, err
		}

		tempTask.SubTasks = append(tempTask.SubTasks, *tempSubTask)

		resTasks = append(resTasks, tempTask)
	}

	taskMap := make(map[int][]entity.Task)

	for _, task := range resTasks {
		taskMap[task.ID] = append(taskMap[task.ID], *task)
	}

	var groupedTasks []*entity.Task

	for _, tasks := range taskMap {
		combinedSubtasks := make([]entity.SubTask, 0)
		for _, task := range tasks {
			if task.SubTasks[0].ID != 0 {
				combinedSubtasks = append(combinedSubtasks, task.SubTasks...)
			}
		}

		baseTask := tasks[0]

		baseTask.SubTasks = combinedSubtasks

		groupedTasks = append(groupedTasks, &baseTask)
	}

	return groupedTasks, nil
}

func (r *TaskRepository) GetCompletedTask() ([]*entity.Task, error) {
	stmt, err := r.db.Prepare(queryGetCompletedTask)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*entity.Task

	for rows.Next() {
		tempTask := new(entity.Task)
		if err := rows.Scan(&tempTask.ID, &tempTask.Title, &tempTask.CreatedAt, &tempTask.Deadline, &tempTask.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, tempTask)
	}

	return tasks, nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	stmt, err := r.db.Prepare(queryDeleteTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) UpdateTask(id int, title string, deadline *time.Time) error {
	stmt, err := r.db.Prepare(queryUpdateTask)

	if err != nil {
		return err
	}
	fmt.Println(deadline)
	if _, err = stmt.Exec(title, deadline, id); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) CompleteTask(id int) error {
	stmt, err := r.db.Prepare(queryCompleteTask)

	if err != nil {
		return err
	}

	completed := time.Now()

	if _, err = stmt.Exec(completed, id); err != nil {
		return err
	}

	return nil
}
