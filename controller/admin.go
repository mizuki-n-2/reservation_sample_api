package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	"github.com/mizuki-n-2/reservation_sample_api/service"
)

type AdminController interface {
	CreateAdmin() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type adminController struct {
	authService     service.AuthService
	adminRepository repository.AdminRepository
}

func NewAdminController(authService service.AuthService, adminRepository repository.AdminRepository) AdminController {
	return &adminController{
		authService: authService, adminRepository: adminRepository,
	}
}

type AdminRequest struct {
	Name     string `json:"name" validate:"omitempty,min=2,max=20"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=30"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AdminID struct {
	AdminID string `json:"admin_id"`
}

func (ac *adminController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AdminRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		admin, err := ac.adminRepository.FindByEmail(req.Email)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		if err = admin.CheckPassword(req.Password); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("パスワードが正しくありません: %w", err))
		}

		token, err := ac.authService.GenerateToken(admin.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := LoginResponse{
			Token: token,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (ac *adminController) CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AdminRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if _, err := ac.adminRepository.FindByEmail(req.Email); err == nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("すでに登録されているメールアドレスです"))
		}

		newAdmin, err := model.NewAdmin(req.Name, req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err = ac.adminRepository.Create(newAdmin); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := AdminID{
			AdminID: newAdmin.ID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
