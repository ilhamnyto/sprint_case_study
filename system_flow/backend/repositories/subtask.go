package repositories

import (
	"database/sql"
	"time"

	"github.com/ilhamnyto/sprint_case_study/entity"
)


var (
	queryCreateSubTask = `
		INSERT INTO subtask (task_id, title, created_at, deadline) VALUES (?, ?, ?, ?)
	`
	queryGetOngoingSubTask = `
		SELECT id, task_id, title, created_at, deadline, completed_at from subtask WHERE completed_at is null
	`
	queryDeleteSubTask = `
		DELETE FROM subtask WHERE id = ?
	`
	queryGetCompletedSubTask = `
		SELECT id, task_id, title, created_at, deadline, completed_at from subtask WHERE completed_at is not null
	`
	queryUpdateSubTask = `
		UPDATE subtask SET title = ? WHERE id = ?
	`
	queryCompleteSubTask = `
		UPDATE subtask SET completed_at = ? WHERE id = ?
	`
)

type InterfaceSubTaskRepository interface {
	Create(subtask *entity.SubTask) error
	GetOngoingTask() ([]*entity.SubTask, error)
	GetCompletedTask() ([]*entity.SubTask, error)
	DeleteTask(id int) (error)
	UpdateTask(id int, title string) error
	CompleteTask(id int) error
}

type SubTaskRepository struct {
	db	*sql.DB
}

func NewSubTaskRepository(db *sql.DB) InterfaceSubTaskRepository {
	return &SubTaskRepository{db: db}
}

func (r *SubTaskRepository) Create(subtask *entity.SubTask) error {
	stmt, err := r.db.Prepare(queryCreateTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(subtask.Title, subtask.CreatedAt, subtask.Deadline); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) GetOngoingTask() ([]*entity.SubTask, error) {
	stmt, err := r.db.Prepare(queryGetOngoingTask)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*entity.SubTask

	for rows.Next() {
		tempTask := new(entity.SubTask)
		if err := rows.Scan(&tempTask.ID, &tempTask.TaskID, &tempTask.Title, &tempTask.CreatedAt, &tempTask.Deadline, &tempTask.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, tempTask)
	}

	return tasks, nil
}

func (r *SubTaskRepository) GetCompletedTask() ([]*entity.SubTask, error) {
	stmt, err := r.db.Prepare(queryGetCompletedTask)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*entity.SubTask

	for rows.Next() {
		tempTask := new(entity.SubTask)
		if err := rows.Scan(&tempTask.ID, &tempTask.TaskID, &tempTask.Title, &tempTask.CreatedAt, &tempTask.Deadline, &tempTask.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, tempTask)
	}

	return tasks, nil
}

func (r *SubTaskRepository) DeleteTask(id int) (error) {
	stmt, err := r.db.Prepare(queryDeleteTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) UpdateTask(id int, title string) error {
	stmt, err := r.db.Prepare(queryUpdateTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(title, id); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) CompleteTask(id int) error {
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

