package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQLConn(connStr string) error {
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	// defer db.Close()
	return db.Ping()
}

func GetDbConnection() (*sql.DB) {
	return db
}
