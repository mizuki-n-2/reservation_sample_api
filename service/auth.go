package service

//go:generate mockgen -source=auth.go -destination=auth_mock.go -package=service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
)

type AuthService interface {
	GenerateToken(adminID string) (string, error)
	ValidateToken(c echo.Context) error
}

type authService struct {
	adminRepository repository.AdminRepository
}

func NewAuthService(adminRepository repository.AdminRepository) AuthService {
	return &authService{adminRepository: adminRepository}
}

type jwtCustomClaims struct {
	AdminID string `json:"admin_id"`
	jwt.StandardClaims
}

func (as *authService) GenerateToken(adminID string) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	claims := &jwtCustomClaims{
		adminID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "reservation_sample",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *authService) ValidateToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("トークンの有効期限が切れています")
	}

	adminID := claims["admin_id"].(string)
	if _, err := as.adminRepository.FindByID(adminID); err != nil {
		return err
	}

	return nil
}
