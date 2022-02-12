package handler

import (
	"github.com/labstack/echo/v4"
)

func GetReservations() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func CreateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(201, "")
	}
}

func GetReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(204, "")
	}
}
