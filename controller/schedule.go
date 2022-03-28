package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/service"
)

type ScheduleController interface {
	CreateSchedule() echo.HandlerFunc
	GetSchedules() echo.HandlerFunc
	GetSchedule() echo.HandlerFunc
	UpdateSchedule() echo.HandlerFunc
	DeleteSchedule() echo.HandlerFunc
}

type scheduleController struct {
	authService        service.AuthService
	scheduleRepository repository.ScheduleRepository
}

func NewScheduleController(
	authService service.AuthService,
	scheduleRepository repository.ScheduleRepository,
) ScheduleController {
	return &scheduleController{
		authService:        authService,
		scheduleRepository: scheduleRepository,
	}
}

type ScheduleRequest struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	MaxNumber int    `json:"max_number"`
}

type ScheduleResponse struct {
	ID                string          `json:"id"`
	Date              model.Date      `json:"date"`
	StartTime         model.StartTime `json:"start_time"`
	ReservationNumber int             `json:"reservation_number"`
	MaxNumber         model.MaxNumber `json:"max_number"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

func (sc *scheduleController) GetSchedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		schedules, err := sc.scheduleRepository.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := make([]ScheduleResponse, len(schedules))
		for i, schedule := range schedules {
			res[i] = ScheduleResponse{
				ID:                schedule.ID,
				Date:              schedule.Date,
				StartTime:         schedule.StartTime,
				ReservationNumber: len(schedule.Reservations),
				MaxNumber:         schedule.MaxNumber,
				CreatedAt:         schedule.CreatedAt,
				UpdatedAt:         schedule.UpdatedAt,
			}
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (sc *scheduleController) GetSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		schedule, err := sc.scheduleRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		res := ScheduleResponse{
			ID:                schedule.ID,
			Date:              schedule.Date,
			StartTime:         schedule.StartTime,
			ReservationNumber: len(schedule.Reservations),
			MaxNumber:         schedule.MaxNumber,
			CreatedAt:         schedule.CreatedAt,
			UpdatedAt:         schedule.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (sc *scheduleController) CreateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := sc.authService.ValidateToken(c); err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		newSchedule, err := model.NewSchedule(req.Date, req.StartTime, req.MaxNumber)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		schedule, err := sc.scheduleRepository.Create(newSchedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := ScheduleResponse{
			ID:                schedule.ID,
			Date:              schedule.Date,
			StartTime:         schedule.StartTime,
			ReservationNumber: 0,
			MaxNumber:         schedule.MaxNumber,
			CreatedAt:         schedule.CreatedAt,
			UpdatedAt:         schedule.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (sc *scheduleController) UpdateSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := sc.authService.ValidateToken(c); err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		id := c.Param("id")
		oldSchedule, err := sc.scheduleRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		err = oldSchedule.UpdateMaxNumber(req.MaxNumber)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		schedule, err := sc.scheduleRepository.Update(&oldSchedule)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := ScheduleResponse{
			ID:                schedule.ID,
			Date:              schedule.Date,
			StartTime:         schedule.StartTime,
			ReservationNumber: len(schedule.Reservations),
			MaxNumber:         schedule.MaxNumber,
			CreatedAt:         schedule.CreatedAt,
			UpdatedAt:         schedule.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (sc *scheduleController) DeleteSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := sc.authService.ValidateToken(c); err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		id := c.Param("id")
		schedule, err := sc.scheduleRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		if err = sc.scheduleRepository.Delete(schedule.ID); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
