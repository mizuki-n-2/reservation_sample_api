package model

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB
var err error

func InitDB() *gorm.DB {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Admin{}, Reservation{}, Schedule{})

	return db
}

func GetDB() *gorm.DB {
	return db
}
