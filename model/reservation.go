package model

import (
	"time"
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
