package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1202174326@tcp(localhost:3306)/todolist")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTaskQuery := `
		CREATE TABLE IF NOT EXISTS task (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255),
			created_at date,
			deadline date,
			completed_at date
		);
	`
	_, err = db.Exec(createTaskQuery)
	if err != nil {
		log.Fatal(err)
	}

	createSubTaskQuery := `
		CREATE TABLE IF NOT EXISTS subtask (
			id INT AUTO_INCREMENT PRIMARY KEY,
			task_id int NOT NULL,
			title VARCHAR(255),
			created_at date,
			deadline date,
			completed_at date,
			FOREIGN KEY (task_id) REFERENCES task(id)
		);
	`
	_, err = db.Exec(createSubTaskQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")
}