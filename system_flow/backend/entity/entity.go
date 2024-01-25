package entity

import "time"


type Task struct {
	ID			int 		`json:"id"`
	Title 		string		`json:"title"`
	CreatedAt 	time.Time	`json:"created_at"`
	Deadline 	*time.Time	`json:"deadline"`
	CompletedAt *time.Time	`json:"completed_at"`
	SubTasks	[]SubTask	`json:"subtasks"`
}

type SubTask struct {
	ID			int 		`json:"id"`
	TaskID		int 		`json:"task_id"`
	Title 		string		`json:"title"`
	CreatedAt 	time.Time	`json:"created_at"`
	Deadline 	*time.Time	`json:"deadline"`
	CompletedAt *time.Time	`json:"completed_at"`
}

