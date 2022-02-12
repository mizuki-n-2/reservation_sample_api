package handler

import (
	"github.com/labstack/echo/v4"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(201, "")
	}
}