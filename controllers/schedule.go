package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/models"
)

type ScheduleRequest struct {
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	MaxNumber int    `json:"max_number"`
}

func GetAvailableSchedules() echo.HandlerFunc {
	return func(c echo.Context) error {
		schedules, err := model.GetAvailableSchedules()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, schedules)
	}
}

func GetAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		scheduleID := c.Param("id")
		schedule, err := model.GetAvailableSchedule(scheduleID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, schedule)
	}
}

func CreateAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminID := getAdminIDFromToken(c)

		if !model.IsAdmin(adminID) {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

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
		adminID := getAdminIDFromToken(c)

		if !model.IsAdmin(adminID) {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		scheduleID := c.Param("id")
		var req ScheduleRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err := model.UpdateAvailableScheduleMaxNumber(scheduleID, req.MaxNumber)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusNoContent, "")
	}
}

func DeleteAvailableSchedule() echo.HandlerFunc {
	return func(c echo.Context) error {
		adminID := getAdminIDFromToken(c)

		if !model.IsAdmin(adminID) {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}
		
		scheduleID := c.Param("id")
		err := model.DeleteAvailableSchedule(scheduleID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusNoContent, "")
	}
}

func getAdminIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	adminID := claims["admin_id"].(string)
	return adminID
}