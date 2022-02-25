package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"regexp"
)

type Reservation struct {
	ID                       string         `json:"id"`
	Name                     Name           `json:"name"`
	Email                    Email          `json:"email"`
	PhoneNumber              PhoneNumber    `json:"phone_number"`
	Address                  Address        `json:"address"`
	AdultNumber              NumberOfPeople `json:"adult_number"`
	PrimarySchoolChildNumber NumberOfPeople `json:"primary_school_child_number"`
	ChildNumber              NumberOfPeople `json:"child_number"`
	SearchID                 string         `json:"search_id"`
	ScheduleID               string         `json:"schedule_id"`
	CreatedAt                time.Time      `json:"created_at"`
	UpdatedAt                time.Time      `json:"updated_at"`
}

func NewReservation(name, email, phoneNumber, address string, adultNumber, primarySchoolChildNumber, childNumber int, scheduleID string) (*Reservation, error) {
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}

	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	newPhoneNumber, err := NewPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	newAddress, err := NewAddress(address)
	if err != nil {
		return nil, err
	}

	newAdultNumber, err := NewNumberOfPeople(adultNumber)
	if err != nil {
		return nil, err
	}

	newPrimarySchoolChildNumber, err := NewNumberOfPeople(primarySchoolChildNumber)
	if err != nil {
		return nil, err
	}

	newChildNumber, err := NewNumberOfPeople(childNumber)
	if err != nil {
		return nil, err
	}

	reservation := &Reservation{
		ID:                       uuid.NewString(),
		Name:                     newName,
		Email:                    newEmail,
		PhoneNumber:              newPhoneNumber,
		Address:                  newAddress,
		AdultNumber:              newAdultNumber,
		PrimarySchoolChildNumber: newPrimarySchoolChildNumber,
		ChildNumber:              newChildNumber,
		SearchID:                 uuid.NewString(),
		ScheduleID:               scheduleID,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	return reservation, nil
}

type PhoneNumber string

func NewPhoneNumber(value string) (PhoneNumber, error) {
	PHONE_NUMBER_PATTERN := `^0\d{1,3}-\d{1,4}-\d{3,4}$`

	if !regexp.MustCompile(PHONE_NUMBER_PATTERN).MatchString(value) {
		return "", fmt.Errorf("電話番号の形式が正しくありません")
	}

	return PhoneNumber(value), nil
}

type Address string

func NewAddress(value string) (Address, error) {
	if len(value) == 0 {
		return "", fmt.Errorf("住所が入力されていません")
	}

	return Address(value), nil
}

type NumberOfPeople int

func NewNumberOfPeople(value int) (NumberOfPeople, error) {
	MAX_NUMBER_OF_PEOPLE := 50
	MIN_NUMBER_OF_PEOPLE := 0

	if value < MIN_NUMBER_OF_PEOPLE || value > MAX_NUMBER_OF_PEOPLE {
		return 0, fmt.Errorf("人数は%d人以上%d人以下にしてください", MIN_NUMBER_OF_PEOPLE, MAX_NUMBER_OF_PEOPLE)
	}

	return NumberOfPeople(value), nil
}
