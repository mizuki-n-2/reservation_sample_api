package infra

import (
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repository.AdminRepository {
	return &adminRepository{db: db}
}

func (ar *adminRepository) Store(admin *model.Admin) (string, error) {
	
}

func (ar *adminRepository) FindByID(id string) (model.Admin, error) {
	
}

func (ar *adminRepository) FindByEmail(email string) (model.Admin, error) {
	
}