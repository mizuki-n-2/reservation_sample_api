package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/mizuki-n-2/reservation_sample_api/models"
)

type CreateReservationRequest struct {
	Name                     string `json:"name"`
	Email                    string `json:"email"`
	PhoneNumber              string `json:"phone_number"`
	Address                  string `json:"address"`
	AdultNumber              int    `json:"adult_number"`
	PrimarySchoolChildNumber int    `json:"primary_school_child_number"`
	ChildNumber              int    `json:"child_number"`
	ScheduleID               string `json:"schedule_id"`
}

func GetReservations() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservations, err := model.GetReservations()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, reservations)
	}
}

func CreateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateReservationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		reservationID, err := model.CreateReservation(req.Name, req.Email, req.PhoneNumber, req.Address, req.AdultNumber, req.PrimarySchoolChildNumber, req.ChildNumber, req.ScheduleID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := map[string]string{
			"id": reservationID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func GetReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservationID := c.Param("id")
		reservation, err := model.GetReservation(reservationID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, reservation)
	}
}

func DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservationID := c.Param("id")
		err := model.DeleteReservation(reservationID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusNoContent, "")
	}
}
