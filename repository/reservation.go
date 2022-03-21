package repository

//go:generate mockgen -source=reservation.go -destination=reservation_mock.go -package=repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type ReservationRepository interface {
	Create(reservation *model.Reservation) (model.Reservation, error)
	FindAll() ([]model.Reservation, error)
	FindByID(id string) (model.Reservation, error)
	Delete(id string) error
}
