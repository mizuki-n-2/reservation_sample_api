package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type ScheduleRequest struct {
	Date string `json:"date"`
	StartTime string `json:"start_time"`
	MaxNumber int `json:"max_number"`
}

func GetAvailableSchedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func GetAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func CreateAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		scheduleID := model.CreateAvailableSchedule(req.Date, req.StartTime, req.MaxNumber)

		res := map[string]string{
			"id": scheduleID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func UpdateAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(204, "")
	}
}

func DeleteAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(204, "")
	}
}