package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
		"fmt"
		  "github.com/joho/godotenv"
		  	"os"
)

var db *gorm.DB

func Connect() {
		_ = godotenv.Load()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, name)

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