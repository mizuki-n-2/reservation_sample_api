package repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type ReservationRepository interface {
	Create(reservation *model.Reservation) (string, error)
	FindAll() ([]model.Reservation, error)
	FindByID(id string) (model.Reservation, error)
	Delete(id string) error
}
