package repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type ScheduleRepository interface {
	Create(schedule *model.Schedule) (model.Schedule, error)
	FindAll() ([]model.Schedule, error)
	FindByID(id string) (model.Schedule, error)
	Update(schedule *model.Schedule) (model.Schedule, error)
	Delete(id string) error
}
