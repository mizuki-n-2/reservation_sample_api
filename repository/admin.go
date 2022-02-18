package repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type AdminRepository interface {
	Create(admin *model.Admin) (string, error)
	FindByID(id string) (model.Admin, error)
	FindByEmail(email string) (model.Admin, error)
}
