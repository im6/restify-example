package store

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// init
func InitMySQLConn() error {
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("SQL_USERNAME"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_PORT"),
		os.Getenv("SQL_DATABASE"),
	)

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
