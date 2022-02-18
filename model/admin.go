package model

import (
	"errors"
	"regexp"
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
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}

	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	
	newPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	admin := &Admin{
		ID:        uuid.NewString(),
		Name:      newName,
		Email:     newEmail,
		Password:  newPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return admin, nil
}

func NewName(value string) (string, error) {
	MIN_LENGTH_USER_NAME := 2
	MAX_LENGTH_USER_NAME := 20

	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_NAME || utf8.RuneCountInString(value) > MAX_LENGTH_USER_NAME {
		return "", fmt.Errorf("nameは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_NAME, MAX_LENGTH_USER_NAME)
	}

	return value, nil
}

func NewEmail(value string) (string, error) {
	EMAIL_PATTERN := `^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	
	if !regexp.MustCompile(EMAIL_PATTERN).MatchString(value) {
		return "", errors.New("emailの形式が正しくありません")
	}

	return value, nil
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
