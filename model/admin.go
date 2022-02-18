package model

import (
	"fmt"
	"unicode/utf8"
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
	newPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	// TODO: 作成時のバリデーション(name, email)
	admin := &Admin{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  newPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return admin, nil
}

func NewPassword(value string) (string, error) {
	MIN_LENGTH_USER_PASSWORD := 8
	MAX_LENGTH_USER_PASSWORD := 30

	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_PASSWORD || utf8.RuneCountInString(value) > MAX_LENGTH_USER_PASSWORD {
		return "", fmt.Errorf("passwordは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_PASSWORD, MAX_LENGTH_USER_PASSWORD)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
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
