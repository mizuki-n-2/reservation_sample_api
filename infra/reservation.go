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

func (rr *reservationRepository) Create(reservation *model.Reservation) (string, error) {
	if err := rr.db.Create(reservation).Error; err != nil {
		return "", err
	}

	return reservation.ID, nil
}

func (rr *reservationRepository) FindAll() ([]model.Reservation, error) {
	var reservations []model.Reservation
	if err := rr.db.Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

func (rr *reservationRepository) FindByID(id string) (model.Reservation, error) {
	var reservation model.Reservation
	if err := rr.db.Where("id = ?", id).First(&reservation).Error; err != nil {
		return model.Reservation{}, err
	}

	return reservation, nil
}

func (rr *reservationRepository) Delete(id string) error {
	if err := rr.db.Where("id = ?", id).Delete(&model.Reservation{}).Error; err != nil {
		return err
	}

	return nil
}
