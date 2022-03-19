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

	// 認証なし
	e.POST("/admins", c.AdminController.CreateAdmin())
	e.POST("/login", c.AdminController.Login())
	e.GET("/reservations", c.ReservationController.GetReservations())
	e.POST("/reservations", c.ReservationController.CreateReservation())	
	e.GET("/reservations/:id", c.ReservationController.GetReservation())	
	e.DELETE("/reservations/:id", c.ReservationController.DeleteReservation())
	e.GET("/schedules", c.ScheduleController.GetSchedules())
	e.GET("/schedules/:id", c.ScheduleController.GetSchedule())

	// 認証あり
	admin := e.Group("/admin")
	admin.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	admin.POST("/schedules", c.ScheduleController.CreateSchedule())
	admin.PATCH("/schedules/:id", c.ScheduleController.UpdateSchedule())
	admin.DELETE("/schedules/:id", c.ScheduleController.DeleteSchedule())

	e.Logger.Fatal(e.Start(":" + port))
}