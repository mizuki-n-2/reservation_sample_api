package main

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizuki-n-2/reservation_sample_api/controllers"
	"github.com/mizuki-n-2/reservation_sample_api/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	model.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	// 認証なし
	e.POST("/signup", controller.CreateAdmin())
	e.POST("/login", controller.Login())
	e.GET("/reservations", controller.GetReservations())
	e.POST("/reservations", controller.CreateReservation())	
	e.GET("/reservations/:id", controller.GetReservation())	
	e.DELETE("/reservations/:id", controller.DeleteReservation())
	e.GET("/available-schedules", controller.GetAvailableSchedules())
	e.GET("/available-schedules/:id", controller.GetAvailableSchedule())

	// 認証あり
	admin := e.Group("/admin")
	admin.Use(middleware.JWT([]byte(os.Getenv("JWT_SIGNING_KEY"))))
	admin.POST("/available-schedules", controller.CreateAvailableSchedule())
	admin.PATCH("/available-schedules/:id", controller.UpdateAvailableSchedule())
	admin.DELETE("/available-schedules/:id", controller.DeleteAvailableSchedule())

	e.Logger.Fatal(e.Start(":8080"))
}