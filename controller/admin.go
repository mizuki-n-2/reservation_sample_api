package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
)

type AdminController interface {
	CreateAdmin() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type adminController struct {
	adminRepository repository.AdminRepository
}

func NewAdminController(adminRepository repository.AdminRepository) AdminController {
	return &adminController{adminRepository: adminRepository}
}

type AdminRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *adminController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AdminRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		admin, err := ac.adminRepository.FindByEmail(req.Email)
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		token, err := admin.Authenticate(req.Password)
		// TODO: エラーハンドリング(パスワードが違う場合とその他で分ける)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := map[string]string{
			"token": token,
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

		newAdmin, err := model.NewAdmin(req.Name, req.Email, req.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdAdminID, err := ac.adminRepository.Create(newAdmin)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := map[string]string{
			"admin_id": createdAdminID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}
