package handler

import (
	"github.com/labstack/echo/v4"
)

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
		return c.JSON(201, "")
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