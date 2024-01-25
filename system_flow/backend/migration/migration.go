package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhamnyto/sprint_case_study/config"
)

func main() {

	config.LoadConfig(".env")

	var (
		db_host = config.GetString(config.DB_HOST)
		db_port = config.GetString(config.DB_PORT)
		db_user = config.GetString(config.DB_USER)
		db_password = config.GetString(config.DB_PASSWORD)
		db_name = config.GetString(config.DB_NAME)
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db_user, db_password, db_host, db_port, db_name,
	)

	db, err := sql.Open("mysql", dsn)

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
			FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE
		);
	`
	_, err = db.Exec(createSubTaskQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully!")
}