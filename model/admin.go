package model

import (
	"os"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAdmin(name, email, password string) (*Admin, error) {
	// TODO: 作成時のバリデーション(name, email, password)
	admin := &Admin{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return admin, nil
}

type MyCustomClaims struct {
	AdminID string `json:"admin_id"`
	jwt.StandardClaims
}

func (admin *Admin) Authenticate(password string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := createToken(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func createToken(adminID string) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	claims := MyCustomClaims{
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
