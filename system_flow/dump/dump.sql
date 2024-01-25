CREATE TABLE IF NOT EXISTS task (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(255),
			created_at date,
			deadline date,
			completed_at date
		);


CREATE TABLE IF NOT EXISTS subtask (
			id INT AUTO_INCREMENT PRIMARY KEY,
			task_id int NOT NULL,
			title VARCHAR(255),
			created_at date,
			deadline date,
			completed_at date,
			FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE
		);