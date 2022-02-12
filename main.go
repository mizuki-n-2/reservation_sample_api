package main

import (
	"os"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizuki-n-2/reservation_sample_api/handler"
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

var db *gorm.DB
var err error

func initDB() *gorm.DB {
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

	db.AutoMigrate(model.Admin{}, model.Reservation{}, model.Schedule{})

	return db
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	db = initDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	// 認証なし
	e.GET("/reservations", handler.GetReservations())
	e.POST("/reservations", handler.CreateReservation())	
	e.GET("/reservations/:id", handler.GetReservation())	
	e.DELETE("/reservations/:id", handler.DeleteReservation())
	e.GET("/available-schedules", handler.GetAvailableSchedules())
	e.GET("/available-schedules/:id", handler.GetAvailableSchedule())

	// 認証あり
	admin := e.Group("/admin")
	admin.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	admin.POST("/", handler.CreateAdmin())
	admin.POST("/login", handler.Login())
	admin.POST("/available-schedules", handler.CreateAvailableSchedule())
	admin.PATCH("/available-schedules/:id", handler.UpdateAvailableSchedule())
	admin.DELETE("/available-schedules/:id", handler.DeleteAvailableSchedule())

	e.Logger.Fatal(e.Start(":8080"))
}