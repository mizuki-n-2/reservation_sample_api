package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/stretchr/testify/assert"
)

func TestAdmin_Login(t *testing.T) {}

func TestAdmin_CreateAdmin(t *testing.T) {
	var ( 
		requestJSON = `{"name":"山田太郎","email":"user@example.com","password":"password123"}`
		responseJSON = `{"admin_id":"22edd7ee-6d0e-416f-a538-25d1bd2cbbf1"}`
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := repository.NewMockAdminRepository(ctrl)
	m.EXPECT().Create(gomock.Any()).Return(nil)

	c := controller.NewAdminController(m)
	c.CreateAdmin()(context)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.IsType(t, responseJSON, rec.Body.String())
}
