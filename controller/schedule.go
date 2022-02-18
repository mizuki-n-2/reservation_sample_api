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
	adminRepository    repository.AdminRepository
}

func NewScheduleController(scheduleRepository repository.ScheduleRepository, adminRepository repository.AdminRepository) ScheduleController {
	return &scheduleController{scheduleRepository: scheduleRepository, adminRepository: adminRepository}
}

type ScheduleRequest struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	MaxNumber int    `json:"max_number"`
}

func (sc *scheduleController) GetSchedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		schedules, err := sc.scheduleRepository.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, schedules)
	}
}

func (sc *scheduleController) GetSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		schedule, err := sc.scheduleRepository.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, schedule)
	}
}

func (sc *scheduleController) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth(c, sc.adminRepository)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

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
		err := auth(c, sc.adminRepository)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		schedule, err := sc.scheduleRepository.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		// TODO: maxNumberのバリデーション
		schedule.MaxNumber = req.MaxNumber

		err = sc.scheduleRepository.Update(&schedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}

func (sc *scheduleController) DeleteSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := auth(c, sc.adminRepository)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		id := c.Param("id")
		_, err = sc.scheduleRepository.FindByID(id)
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

func auth(c echo.Context, adminRepository repository.AdminRepository) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	adminID := claims["admin_id"].(string)

	_, err := adminRepository.FindByID(adminID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "admin_id is invalid")
	}

	return nil
}
