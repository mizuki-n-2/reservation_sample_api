package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"net/http"
)

type CreateAdminRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateAdminRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		token, err := model.Login(req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := map[string]string{
			"token": token,
		}

		return c.JSON(200, res)
	}
}

func CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateAdminRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		adminID, err := model.CreateAdmin(req.Name, req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := map[string]string{
			"id": adminID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
