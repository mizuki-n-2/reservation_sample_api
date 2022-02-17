package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
)

type ReservationController interface {
	CreateReservation() echo.HandlerFunc
	GetReservations() echo.HandlerFunc
	GetReservation() echo.HandlerFunc
	DeleteReservation() echo.HandlerFunc
}

type reservationController struct {
	reservationRepository repository.ReservationRepository
}

func NewReservationController(reservationRepository repository.ReservationRepository) ReservationController {
	return &reservationController{reservationRepository: reservationRepository}
}

type ReservationRequest struct {
	Name                     string `json:"name"`
	Email                    string `json:"email"`
	PhoneNumber              string `json:"phone_number"`
	Address                  string `json:"address"`
	AdultNumber              int    `json:"adult_number"`
	PrimarySchoolChildNumber int    `json:"primary_school_child_number"`
	ChildNumber              int    `json:"child_number"`
	ScheduleID               string `json:"schedule_id"`
}

func (rc *reservationController) GetReservations() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (rc *reservationController) CreateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO	
	}
}

func (rc *reservationController) GetReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}

func (rc *reservationController) DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
	}
}
