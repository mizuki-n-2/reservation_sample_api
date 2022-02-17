package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
)

type ScheduleController interface {
	CreateSchedule() echo.HandlerFunc
	GetSchedules() echo.HandlerFunc
	GetSchedule() echo.HandlerFunc
	UpdateSchedule() echo.HandlerFunc
	DeleteSchedule() echo.HandlerFunc
}

type scheduleController struct {
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleController(scheduleRepository repository.ScheduleRepository) ScheduleController {
	return &scheduleController{scheduleRepository: scheduleRepository}
}

type ScheduleRequest struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	MaxNumber int    `json:"max_number"`
}

func (sc *scheduleController) GetSchedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (sc *scheduleController) GetSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (sc *scheduleController) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (sc *scheduleController) UpdateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (sc *scheduleController) DeleteSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

// これをどこに置くか
func getAdminIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	adminID := claims["admin_id"].(string)
	return adminID
}