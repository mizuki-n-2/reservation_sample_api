package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"net/http"
)

type ReservationController interface {
	CreateReservation() echo.HandlerFunc
	GetReservations() echo.HandlerFunc
	GetReservation() echo.HandlerFunc
	DeleteReservation() echo.HandlerFunc
}

type reservationController struct {
	reservationRepository repository.ReservationRepository
	scheduleRepository    repository.ScheduleRepository
}

func NewReservationController(reservationRepository repository.ReservationRepository, scheduleRepository repository.ScheduleRepository) ReservationController {
	return &reservationController{reservationRepository: reservationRepository, scheduleRepository: scheduleRepository}
}

type ReservationRequest struct {
	Name                     string `json:"name" validate:"required,min=2,max=20"`
	Email                    string `json:"email" validate:"required,email"`
	PhoneNumber              string `json:"phone_number" validate:"required,max=14"`
	Address                  string `json:"address" validate:"required,max=50"`
	AdultNumber              int    `json:"adult_number" validate:"gte=0,lte=50"`
	PrimarySchoolChildNumber int    `json:"primary_school_child_number" validate:"gte=0,lte=50"`
	ChildNumber              int    `json:"child_number" validate:"gte=0,lte=50"`
	ScheduleID               string `json:"schedule_id" validate:"required,uuid4"`
}

func (rc *reservationController) GetReservations() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservations, err := rc.reservationRepository.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, reservations)
	}
}

func (rc *reservationController) CreateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ReservationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		newReservation, err := model.NewReservation(req.Name, req.Email, req.PhoneNumber, req.Address, req.AdultNumber, req.PrimarySchoolChildNumber, req.ChildNumber, req.ScheduleID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		_, err = rc.scheduleRepository.FindByID(req.ScheduleID)
		// TODO: 適切なエラーハンドリング(スケジュールが存在しない場合)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		reservation, err := rc.reservationRepository.Create(newReservation)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, reservation)
	}
}

func (rc *reservationController) GetReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		reservation, err := rc.reservationRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, reservation)
	}
}

func (rc *reservationController) DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		reservation, err := rc.reservationRepository.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		err = rc.reservationRepository.Delete(reservation.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
