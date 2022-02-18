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

		reservation, err := model.NewReservation(req.Name, req.Email, req.PhoneNumber, req.Address, req.AdultNumber, req.PrimarySchoolChildNumber, req.ChildNumber, req.ScheduleID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		_, err = rc.scheduleRepository.FindByID(req.ScheduleID)
		// TODO: 適切なエラーハンドリング(スケジュールが存在しない場合)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		reservationID, err := rc.reservationRepository.Create(reservation)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := map[string]string{
			"reservation_id": reservationID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (rc *reservationController) GetReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		reservation, err := rc.reservationRepository.FindByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, reservation)
	}
}

func (rc *reservationController) DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := rc.reservationRepository.Delete(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
