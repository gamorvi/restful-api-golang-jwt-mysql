package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB
var err error
var e error

func init() {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	charset := os.Getenv("charset")
	parseTime := os.Getenv("parse_time")

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&&parseTime=%s", username, password, dbHost, dbPort, dbName, charset, parseTime)
	conn, err := gorm.Open(dbType, dbUri)

	db = conn
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database on " + dbHost)

}

func GetDB() *gorm.DB {
	return db
}
