package database

import (
	"database/sql"
	"fmt"

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
		db_port = config.GetString(config.DB_PORT)
		db_user = config.GetString(config.DB_USER)
		db_password = config.GetString(config.DB_PASSWORD)
		db_name = config.GetString(config.DB_NAME)
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db_user, db_password, db_host, db_port, db_name,
	)

	dbsql, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	if err := dbsql.Ping(); err != nil {
		panic(err)
	}

	db := NewDatabase()
	db = db.SetDatabase(dbsql)
	
	return db
}
