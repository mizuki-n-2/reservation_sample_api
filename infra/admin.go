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

func (ar *adminRepository) Create(admin *model.Admin) (string, error) {
	if err := ar.db.Create(admin).Error; err != nil {
		return "", err
	}

	return admin.ID, nil
}

func (ar *adminRepository) FindByID(id string) (model.Admin, error) {
	var admin model.Admin
	if err := ar.db.Where("id = ?", id).First(&admin).Error; err != nil {
		return model.Admin{}, err
	}

	return admin, nil
}

func (ar *adminRepository) FindByEmail(email string) (model.Admin, error) {
	var admin model.Admin
	if err := ar.db.Find(&admin, model.Admin{Email: email}).Error; err != nil {
		return model.Admin{}, err
	}

	return admin, nil
}
