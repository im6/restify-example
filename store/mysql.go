package store

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql include
)

// DB is the db connection instance
var DB *gorm.DB

// ConnectMySQL init conn
func ConnectMySQL() {
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("SQL_USERNAME"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_DATABASE"),
	)

	db, err := gorm.Open("mysql", conn)
	// defer db.Close()

	if err != nil {
		panic(err)
	}

	DB = db

}
