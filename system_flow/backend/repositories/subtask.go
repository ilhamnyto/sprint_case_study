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
		UPDATE subtask SET title = ?, deadline = ? WHERE id = ?
	`
	queryCompleteSubTask = `
		UPDATE subtask SET completed_at = ? WHERE id = ?
	`
)

type InterfaceSubTaskRepository interface {
	CreateSubTask(subtask *entity.SubTask) error
	GetOngoingSubTask() ([]*entity.SubTask, error)
	GetCompletedSubTask() ([]*entity.SubTask, error)
	DeleteSubTask(id int) (error)
	UpdateSubTask(id int, title string, deadline *time.Time) error
	CompleteSubTask(id int) error
}

type SubTaskRepository struct {
	db	*sql.DB
}

func NewSubTaskRepository(db *sql.DB) InterfaceSubTaskRepository {
	return &SubTaskRepository{db: db}
}

func (r *SubTaskRepository) CreateSubTask(subtask *entity.SubTask) error {
	stmt, err := r.db.Prepare(queryCreateSubTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(subtask.TaskID, subtask.Title, subtask.CreatedAt, subtask.Deadline); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) GetOngoingSubTask() ([]*entity.SubTask, error) {
	stmt, err := r.db.Prepare(queryGetOngoingSubTask)

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

func (r *SubTaskRepository) GetCompletedSubTask() ([]*entity.SubTask, error) {
	stmt, err := r.db.Prepare(queryGetCompletedSubTask)

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

func (r *SubTaskRepository) DeleteSubTask(id int) (error) {
	stmt, err := r.db.Prepare(queryDeleteSubTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) UpdateSubTask(id int, title string, deadline *time.Time) error {
	stmt, err := r.db.Prepare(queryUpdateSubTask)

	if err != nil {
		return err
	}

	if _, err = stmt.Exec(title, deadline, id); err != nil {
		return err
	}

	return nil
}

func (r *SubTaskRepository) CompleteSubTask(id int) error {
	stmt, err := r.db.Prepare(queryCompleteSubTask)

	if err != nil {
		return err
	}

	completed := time.Now()

	if _, err = stmt.Exec(completed, id); err != nil {
		return err
	}

	return nil
}

