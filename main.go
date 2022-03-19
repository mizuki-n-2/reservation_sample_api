package main

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizuki-n-2/reservation_sample_api/di"
	"github.com/mizuki-n-2/reservation_sample_api/router"
)

func initDB() *gorm.DB {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	port := os.Getenv("PORT")

	db := initDB()

	c := di.InitDI(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.NewRouter(e, c)

	e.Logger.Fatal(e.Start(":" + port))
}