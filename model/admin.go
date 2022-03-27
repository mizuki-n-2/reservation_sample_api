package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
	"unicode/utf8"
)

type Admin struct {
	ID        string    `json:"id"`
	Name      Name      `json:"name"`
	Email     Email     `json:"email"`
	Password  Password  `json:"password"`
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

type Name string

var (
	MIN_LENGTH_USER_NAME = 2
	MAX_LENGTH_USER_NAME = 20
)

func NewName(value string) (Name, error) {
	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_NAME || utf8.RuneCountInString(value) > MAX_LENGTH_USER_NAME {
		return "", fmt.Errorf("nameは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_NAME, MAX_LENGTH_USER_NAME)
	}

	return Name(value), nil
}

type Email string

var (
	EMAIL_PATTERN = `^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
)

func NewEmail(value string) (Email, error) {
	if !regexp.MustCompile(EMAIL_PATTERN).MatchString(value) {
		return "", errors.New("emailの形式が正しくありません")
	}

	return Email(value), nil
}

type Password string

var (
	MIN_LENGTH_USER_PASSWORD = 8
	MAX_LENGTH_USER_PASSWORD = 30
)

func NewPassword(value string) (Password, error) {
	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_PASSWORD || utf8.RuneCountInString(value) > MAX_LENGTH_USER_PASSWORD {
		return "", fmt.Errorf("passwordは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_PASSWORD, MAX_LENGTH_USER_PASSWORD)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(hashedPassword), nil
}

func (admin *Admin) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
