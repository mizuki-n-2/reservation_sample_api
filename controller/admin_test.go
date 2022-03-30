package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/service"
	v "github.com/mizuki-n-2/reservation_sample_api/validator"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
)

func TestAdmin_Login(t *testing.T) {
	var (
		request = controller.AdminRequest{
			Email: "user@example.com",
			Password: "password123",
		}
		token = "created-token"
	)

	hashedPassword, err := model.NewPassword(request.Password)
	if err != nil {
		t.Fatal(err)
	}

	admin := model.Admin{
		ID: "admin-id",
		Name: "admin-name",
		Email: "user@example.com",
		Password: hashedPassword,
	}

	e := echo.New()
	e.Validator = &v.CustomValidator{Validator: validator.New()}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/admin/login", strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m1.EXPECT().GenerateToken(admin.ID).Return(token, nil)
	m2 := repository.NewMockAdminRepository(ctrl)
	m2.EXPECT().FindByEmail(request.Email).Return(&admin, nil)
	
	c := controller.NewAdminController(m1, m2)
	if err := c.Login()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var response controller.LoginResponse
	if err = json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Token != token {
		t.Fatal("token not match")
	}
}

func TestAdmin_CreateAdmin(t *testing.T) {
	var ( 
		request = controller.AdminRequest{
			Name:     "山田太郎",
			Email:    "user@example.com",
			Password: "password123",
		}
	)

	e := echo.New()
	e.Validator = &v.CustomValidator{Validator: validator.New()}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/admin", strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m2 := repository.NewMockAdminRepository(ctrl)
	m2.EXPECT().FindByEmail(request.Email).Return(nil, gorm.ErrRecordNotFound)
	m2.EXPECT().Create(gomock.Any()).Return(nil)

	c := controller.NewAdminController(m1, m2)
	if err := c.CreateAdmin()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)
	
	var response controller.AdminID
	if err = json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.AdminID == "" {
		t.Fatal("admin_id not exists")
	}
}
