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
		// adminID := getAdminIDFromToken(c)
		schedules, err := sc.scheduleRepository.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, schedules)
	}
}

func (sc *scheduleController) GetSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// adminID := getAdminIDFromToken(c)
		schedule, err := sc.scheduleRepository.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, schedule)
	}
}

func (sc *scheduleController) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		schedule, err := model.NewSchedule(req.Date, req.StartTime, req.MaxNumber)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		scheduleID, err := sc.scheduleRepository.Create(schedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := map[string]string{
			"schedule_id": scheduleID,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (sc *scheduleController) UpdateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// adminID := getAdminIDFromToken(c)
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		schedule, err := sc.scheduleRepository.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		err = sc.scheduleRepository.Update(&schedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (sc *scheduleController) DeleteSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		// adminID := getAdminIDFromToken(c)
		id := c.Param("id")
		_, err := sc.scheduleRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		err = sc.scheduleRepository.Delete(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

// これをどこに置くか
func getAdminIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	adminID := claims["admin_id"].(string)
	return adminID
}