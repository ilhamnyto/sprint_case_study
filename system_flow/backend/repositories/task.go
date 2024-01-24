package repositories

import (
	"database/sql"
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
	queryUpdateTask = `
		UPDATE task SET title = ? WHERE id = ?
	`
	queryCompleteTask = `
		UPDATE task SET completed_at = ? WHERE id = ?
	`
)

type InterfaceTaskRepository interface {
	Create(task *entity.Task) error
	GetOngoingTask() ([]*entity.Task, error)
	GetCompletedTask() ([]*entity.Task, error)
	DeleteTask(id int) error
	UpdateTask(id int, title string) error
	CompleteTask(id int) error
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) InterfaceTaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *entity.Task) error {
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

func (r *TaskRepository) UpdateTask(id int, title string) error {
	stmt, err := r.db.Prepare(queryUpdateTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(title, id); err != nil {
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
