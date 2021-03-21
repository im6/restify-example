package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func getSqlConn() string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configuration.SQL_USERNAME,
		configuration.SQL_PASSWORD,
		configuration.SQL_HOST,
		configuration.SQL_PORT,
		configuration.SQL_DATABASE,
	)
	return connStr
}

func InitMySQLConn() error {
	var err error
	db, err = sql.Open("mysql", getSqlConn())
	if err != nil {
		return err
	}
	// defer db.Close()
	return db.Ping()
}

func GetDbConnection() *sql.DB {
	return db
}
