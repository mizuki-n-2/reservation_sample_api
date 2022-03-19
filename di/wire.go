//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/infra"
	"gorm.io/gorm"
)

type DIContainer struct {
	AdminController       controller.AdminController
	ScheduleController    controller.ScheduleController
	ReservationController controller.ReservationController
}

func InitDI(db *gorm.DB) *DIContainer {
	wire.Build(
		infra.NewAdminRepository,
		infra.NewScheduleRepository,
		infra.NewReservationRepository,
		controller.NewAdminController,
		controller.NewScheduleController,
		controller.NewReservationController,
		wire.Struct(new(DIContainer), "*"),
	)
	return &DIContainer{}
}
