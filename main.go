package main

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/infra"
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

	db := initDB()

	// DI
	adminRepository := infra.NewAdminRepository(db)
	scheduleRepository := infra.NewScheduleRepository(db)
	reservationRepository := infra.NewReservationRepository(db)
	adminController := controller.NewAdminController(adminRepository)
	scheduleController := controller.NewScheduleController(scheduleRepository, adminRepository)
	reservationController := controller.NewReservationController(reservationRepository, scheduleRepository)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 認証なし
	e.POST("/admins", adminController.CreateAdmin())
	e.POST("/login", adminController.Login())
	e.GET("/reservations", reservationController.GetReservations())
	e.POST("/reservations", reservationController.CreateReservation())	
	e.GET("/reservations/:id", reservationController.GetReservation())	
	e.DELETE("/reservations/:id", reservationController.DeleteReservation())
	e.GET("/schedules", scheduleController.GetSchedules())
	e.GET("/schedules/:id", scheduleController.GetSchedule())

	// 認証あり
	admin := e.Group("/admin")
	admin.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	admin.POST("/schedules", scheduleController.CreateSchedule())
	admin.PATCH("/schedules/:id", scheduleController.UpdateSchedule())
	admin.DELETE("/schedules/:id", scheduleController.DeleteSchedule())

	e.Logger.Fatal(e.Start(":8080"))
}