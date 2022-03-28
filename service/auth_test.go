package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/service"
	"github.com/stretchr/testify/assert"
)

func TestAuth_GenerateToken(t *testing.T) {
	var (
		adminID = "admin-id"
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := repository.NewMockAdminRepository(ctrl)

	s := service.NewAuthService(m)
	token, err := s.GenerateToken(adminID)
	
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
