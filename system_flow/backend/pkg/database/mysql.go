package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhamnyto/sprint_case_study/config"
)

type Database struct {
	DbSQL *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) SetDatabase(db *sql.DB) *Database {
	d.DbSQL = db
	return d
}

func ConnectDB() (*Database) {
	var (
		db_host = config.GetString(config.DB_HOST)
		// db_port = config.GetString(config.DB_PORT)
		db_user = config.GetString(config.DB_USER)
		db_password = config.GetString(config.DB_PASSWORD)
		db_name = config.GetString(config.DB_NAME)
		dbsql *sql.DB
		err error
	)

	maxRetries := 10
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?parseTime=true",
			db_user, db_password, db_host, db_name,
		)
	
		dbsql, err = sql.Open("mysql", dsn)
		
		if err := dbsql.Ping(); err == nil {
			break
		}

		log.Printf("Error connecting to MySQL: %v", err)
		log.Printf("Retrying in %s...", retryInterval)
		time.Sleep(retryInterval)
	}



	db := NewDatabase()
	db = db.SetDatabase(dbsql)
	
	return db
}
