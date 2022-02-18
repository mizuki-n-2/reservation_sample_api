package model

import (
	"time"
	"github.com/google/uuid"
)

type Reservation struct {
	ID                       string    `json:"id"`
	Name                     string    `json:"name"`
	Email                    string    `json:"email"`
	PhoneNumber              string    `json:"phone_number"`
	Address                  string    `json:"address"`
	AdultNumber              int       `json:"adult_number"`
	PrimarySchoolChildNumber int       `json:"primary_school_child_number"`
	ChildNumber              int       `json:"child_number"`
	SearchID                 string    `json:"search_id"`
	ScheduleID               string    `json:"schedule_id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func NewReservation(name, email, phoneNumber, address string, adultNumber, primarySchoolChildNumber, childNumber int, scheduleID string) (*Reservation, error) {
	// TODO: 作成時の(引数の)バリデーション
	reservation := &Reservation{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		PhoneNumber: phoneNumber,
		Address:   address,
		AdultNumber: adultNumber,
		PrimarySchoolChildNumber: primarySchoolChildNumber,
		ChildNumber: childNumber,
		SearchID: uuid.NewString(),
		ScheduleID: scheduleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return reservation, nil
}
