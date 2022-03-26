package repository

//go:generate mockgen -source=admin.go -destination=admin_mock.go -package=repository

import (
	"github.com/mizuki-n-2/reservation_sample_api/model"
)

type AdminRepository interface {
	Create(admin *model.Admin) error
	FindByID(id string) (model.Admin, error)
	FindByEmail(email string) (model.Admin, error)
}
