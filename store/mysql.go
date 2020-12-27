package store

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the db connection instance
var DB *gorm.DB

// ConnectMySQL init conn
func ConnectMySQL() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("SQL_USERNAME"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_PORT"),
		os.Getenv("SQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	// defer db.Close()

	if err != nil {
		panic(err)
	}

	DB = db

}
