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
	"github.com/stretchr/testify/assert"
)

func TestAdmin_Login(t *testing.T) {
	var (
		request = controller.AdminRequest{
			Email: "user@example.com",
			Password: "password123",
		}
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
	m := repository.NewMockAdminRepository(ctrl)
	m.EXPECT().FindByEmail(request.Email).Return(admin, nil)
	
	c := controller.NewAdminController(m)
	if err := c.Login()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var response controller.LoginResponse
	if err = json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Token == "" {
		t.Fatal("token not exists")
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
	m := repository.NewMockAdminRepository(ctrl)
	m.EXPECT().Create(gomock.Any()).Return(nil)

	c := controller.NewAdminController(m)
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
