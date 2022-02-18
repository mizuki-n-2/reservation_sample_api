package infra

import (
	"log"

	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) repository.ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (sr *scheduleRepository) Create(schedule *model.Schedule) (string, error) {
	if err := sr.db.Create(schedule).Error; err != nil {
		return "", err
	}

	return schedule.ID, nil
}

func (sr *scheduleRepository) Update(schedule *model.Schedule) error {
	if err := sr.db.Save(schedule).Error; err != nil {
		return err
	}

	return nil
}

func (sr *scheduleRepository) FindAll() ([]model.Schedule, error) {
	var schedules []model.Schedule
	var reservations []model.Reservation
	err := sr.db.Model(&schedules).Association("Reservations").Find(&reservations)
	if err != nil {
		return nil, err
	}

	log.Println(reservations)
	log.Println(schedules)

	return schedules, nil
}

func (sr *scheduleRepository) FindByID(id string) (model.Schedule, error) {
	var schedule model.Schedule
	if err := sr.db.Where("id = ?", id).First(&schedule).Error; err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

func (sr *scheduleRepository) Delete(id string) error {
	if err := sr.db.Where("id = ?", id).Delete(&model.Schedule{}).Error; err != nil {
		return err
	}

	return nil
}
