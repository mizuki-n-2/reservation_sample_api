//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/infra"
	"github.com/mizuki-n-2/reservation_sample_api/service"
	"gorm.io/gorm"
)

type Controllers struct {
	AdminController       controller.AdminController
	ScheduleController    controller.ScheduleController
	ReservationController controller.ReservationController
}

func InitDI(db *gorm.DB) *Controllers {
	wire.Build(
		infra.NewAdminRepository,
		infra.NewScheduleRepository,
		infra.NewReservationRepository,
		service.NewAuthService,
		controller.NewAdminController,
		controller.NewScheduleController,
		controller.NewReservationController,
		wire.Struct(new(Controllers), "*"),
	)
	return &Controllers{}
}
