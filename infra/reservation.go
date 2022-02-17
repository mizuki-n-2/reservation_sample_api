package infra

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) repository.ReservationRepository {
	return &reservationRepository{db: db}
}

func (rr *reservationRepository) Store(reservation *model.Reservation) (string, error) {

}

func (rr *reservationRepository) FindAll() ([]model.Reservation, error) {

}

func (rr *reservationRepository) FindByID(id string) (model.Reservation, error) {

}

func (rr *reservationRepository) Delete(id string) error {

}
