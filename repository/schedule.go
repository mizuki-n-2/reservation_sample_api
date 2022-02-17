package repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type ScheduleRepository interface {
	Store(schedule *model.Schedule) (string, error)
	FindAll() ([]model.Schedule, error)
	FindByID(id string) (model.Schedule, error)
	Delete(id string) error
}
