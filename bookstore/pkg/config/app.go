package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
		"fmt"
)

var db *gorm.DB

func Connect() {
	dsn := "go_user:YourPassword123!@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to DB: %v", err))
	}

	fmt.Println("Connection to DB established successfully")
	db = d
}


// GetDB returns the current active database connection
// This function is used by other packages to access the DB connection
func GetDB() *gorm.DB{
	return db
}