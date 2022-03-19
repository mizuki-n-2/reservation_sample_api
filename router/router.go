package router

import (
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mizuki-n-2/reservation_sample_api/di"
)

func NewRouter(e *echo.Echo, c *di.Controllers) {
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
}