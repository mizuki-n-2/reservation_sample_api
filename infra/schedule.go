package infra

import (
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) repository.ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (sr *scheduleRepository) Store(schedule *model.Schedule) (string, error) {
	
}

func (sr *scheduleRepository) FindAll() ([]model.Schedule, error) {
	
}

func (sr *scheduleRepository) FindByID(id string) (model.Schedule, error) {
	
}

func (sr *scheduleRepository) Delete(id string) error {
	
}