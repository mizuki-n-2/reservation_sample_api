package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID        string    `json:"id" gorm:"primaryKey;size:36"`
	Name      string    `json:"name" gorm:"not null;size:20"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}

type MyCustomClaims struct {
	AdminID string `json:"admin_id"`
	jwt.StandardClaims
}

func CreateAdmin(name, email, password string) (string, error) {
	db := GetDB()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	admin := Admin{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
	}

	if err := db.Create(&admin).Error; err != nil {
		return "", err
	}

	return admin.ID, nil
}

func Login(email, password string) (string, error) {
	db := GetDB()

	admin := Admin{}
	db.Where("email = ?", email).First(&admin)
	if err != nil {
		return "", err
	}

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